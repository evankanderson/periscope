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

package remoteproxy

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"
	"sync"

	"github.com/evankanderson/periscope/pkg/periscope"
	"google.golang.org/grpc"
)

type LocalProxy struct {
	periscope.UnimplementedPeriscopeServer
	httpServer *http.Server
	stream     periscope.Periscope_OutServer

	grpcAddr string

	// This contains the set of outstanding locally-proxied requests awaiting
	// responses over the (singular) grpc stream.
	awaiting map[int64]chan *periscope.ProxyResponse
	lock     sync.Mutex
}

func NewLocalProxy(httpPort int, grpcPort int) (*LocalProxy, error) {

	ret := LocalProxy{
		httpServer: &http.Server{
			Addr: fmt.Sprintf(":%d", httpPort),
		},
		grpcAddr: fmt.Sprintf(":%d", grpcPort),

		awaiting: make(map[int64]chan *periscope.ProxyResponse),
		lock:     sync.Mutex{},
	}
	ret.httpServer.Handler = &ret
	return &ret, nil
}

func (s *LocalProxy) Start() error {
	lis, err := net.Listen("tcp", s.grpcAddr)
	if err != nil {
		return err
	}
	grpc := grpc.NewServer()
	periscope.RegisterPeriscopeServer(grpc, s)
	go grpc.Serve(lis)
	defer grpc.Stop()
	if err := s.httpServer.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func (s *LocalProxy) In(ctx context.Context, in *periscope.ProxyRequest) (*periscope.ProxyResponse, error) {
	req, err := periscope.ReqToHttp(in)
	log.Printf("IN: %s", req.URL)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return periscope.HttpToResp(*resp)
}

func (s *LocalProxy) Out(stream periscope.Periscope_OutServer) error {
	func() {
		s.lock.Lock()
		defer s.lock.Unlock()
		s.stream = stream
	}()
	defer func() {
		s.lock.Lock()
		defer s.lock.Unlock()
		s.stream = nil
	}()

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		func() {
			s.lock.Lock()
			defer s.lock.Unlock()
			if c := s.awaiting[in.Id]; c != nil {
				c <- in
				close(c)
				delete(s.awaiting, in.Id)
				log.Printf("REV DONE %d: %s", in.Id, in.Reason)
			}
		}()
	}
}

func (s *LocalProxy) proxyRequest(in *periscope.ProxyRequest, stream periscope.Periscope_OutServer) {
	return
}

func (s *LocalProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	send, err := periscope.HttpToReq(*r)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	send.Id = rand.Int63()
	c := make(chan (*periscope.ProxyResponse), 1)
	var stream periscope.Periscope_OutServer
	func() {
		s.lock.Lock()
		defer s.lock.Unlock()
		s.awaiting[send.Id] = c
		stream = s.stream
	}()
	log.Printf("REV %d: %s", send.Id, r.URL)
	stream.Send(send)

	out := <-c

	for k, v := range out.Headers {
		w.Header().Set(k, v)
	}
	w.WriteHeader(int(out.Status))
	w.Write(out.Body)
}
