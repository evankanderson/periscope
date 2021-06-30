/*
Copyright © 2021 Evan Anderson <Evan.K.Anderson@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package localproxy

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/elazarl/goproxy"
	"github.com/evankanderson/periscope/pkg/periscope"
	"google.golang.org/grpc"
)

func StartLocalProxy(port int, localTarget string, server string) error {
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	client := periscope.NewPeriscopeClient(conn)

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	proxy.OnRequest().DoFunc(forward(client))
	listenAddr := fmt.Sprintf("localhost:%d", port)
	log.Printf("Listening on %q, forwarding to %q. Incoming will connect to %q", listenAddr, server, localTarget)
	httpServer := &http.Server{
		Addr:    listenAddr,
		Handler: proxy,
	}

	go httpServer.ListenAndServe()
	defer httpServer.Shutdown(context.Background())
	return startReverse(client, localTarget)
}

func forward(client periscope.PeriscopeClient) func(*http.Request, *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	return func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		localError := func(message string, err error) (*http.Request, *http.Response) {
			return r, &http.Response{
				StatusCode: 500,
				Status:     fmt.Sprintf("%s: %s", message, err),
			}
		}

		send, err := periscope.HttpToReq(*r)
		if err != nil {
			return localError("Failed encode", err)
		}
		out, err := client.In(context.Background(), send)
		if err != nil {
			return localError("Failed request", err)
		}
		resp, err := periscope.RespToHttp(out)
		if err != nil {
			return localError("Failed decode", err)
		}
		return r, resp
	}
}

func startReverse(client periscope.PeriscopeClient, localTarget string) error {
	localDial := func(ctx context.Context, network, addr string) (net.Conn, error) {
		if network != "tcp" {
			return nil, fmt.Errorf("Unsupported protocol %q", network)
		}
		return net.Dial("tcp", localTarget)
	}
	httpClient := http.Client{
		Transport: &http.Transport{
			DialContext: localDial,
		},
	}

	stream, err := client.Out(context.Background())
	if err != nil {
		return err
	}
	if err := stream.Send(&periscope.ProxyResponse{
		Id:      -1,
		Status:  100,
		Reason:  "Started",
		Headers: map[string]string{"Preflight": "true"},
		Body:    []byte{},
	}); err != nil {
		return err
	}
	for {
		in, err := stream.Recv()
		if err != nil {
			return err
		}
		go localRequest(in, httpClient, stream)
	}
	// return nil
}

func localRequest(in *periscope.ProxyRequest, client http.Client, stream periscope.Periscope_OutClient) {
	errorResponse := func(message string, err error) {
		stream.Send(&periscope.ProxyResponse{
			Id:     in.Id,
			Status: 500,
			Reason: fmt.Sprintf("%s: %s", message, err),
		})
	}

	req, err := periscope.ReqToHttp(in)
	if err != nil {
		errorResponse("Failed encode", err)
		return
	}
	log.Printf("LOCAL: %s", req.URL)

	resp, err := client.Do(req)
	if err != nil {
		errorResponse("Failed request", err)
		return
	}
	out, err := periscope.HttpToResp(*resp)
	if err != nil {
		errorResponse("Failed decode", err)
		return
	}
	out.Id = in.Id
	log.Printf("LOCAL resp: %s", out.Reason)
	if err := stream.Send(out); err != nil {
		log.Printf("Failed to stream response: %s", err)
	}
}
