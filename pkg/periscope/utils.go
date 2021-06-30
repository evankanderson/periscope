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

package periscope

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func ReqToHttp(in *ProxyRequest) (*http.Request, error) {
	url, err := url.Parse(in.Target)
	if err != nil {
		return nil, err
	}
	url.Host = in.Host
	headers := make(http.Header, len(in.Headers))
	for k, v := range in.Headers {
		headers.Set(k, v)
	}
	return &http.Request{
		Method: in.Verb,
		URL:    url,
		Header: headers,
		Body:   io.NopCloser(bytes.NewReader(in.Body)),
	}, nil
}

func RespToHttp(in *ProxyResponse) (*http.Response, error) {
	headers := make(http.Header, len(in.Headers))
	for k, v := range in.Headers {
		headers.Set(k, v)
	}
	return &http.Response{
		StatusCode: int(in.Status),
		Status:     in.Reason,
		Header:     headers,
		Body:       io.NopCloser(bytes.NewReader(in.Body)),
	}, nil
}

// Closes in.Body
func HttpToResp(in http.Response) (*ProxyResponse, error) {
	defer in.Body.Close()
	body, err := io.ReadAll(in.Body)
	if err != nil {
		return nil, err
	}
	headers := make(map[string]string, len(in.Header))
	for k, v := range in.Header {
		// https://datatracker.ietf.org/doc/html/rfc7230#section-3.2.2
		headers[k] = strings.Join(v, ",")
	}
	return &ProxyResponse{
		Status:  int32(in.StatusCode),
		Reason:  in.Status,
		Headers: headers,
		Body:    body,
	}, nil
}

// Closes in.Body
func HttpToReq(in http.Request) (*ProxyRequest, error) {
	defer in.Body.Close()
	body, err := io.ReadAll(in.Body)
	if err != nil {
		return nil, err
	}
	headers := make(map[string]string, len(in.Header))
	for k, v := range in.Header {
		// https://datatracker.ietf.org/doc/html/rfc7230#section-3.2.2
		headers[k] = strings.Join(v, ",")
	}
	return &ProxyRequest{
		Verb:    in.Method,
		Target:  in.RequestURI,
		Host:    in.Host,
		Headers: headers,
		Body:    body,
	}, nil
}
