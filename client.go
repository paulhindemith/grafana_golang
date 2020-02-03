/*
	Copyright 2016 Alexander I.Grafov <grafov@gmail.com>
	Copyright 2016-2019 The Grafana SDK authors

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

	  http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.

	ॐ तारे तुत्तारे तुरे स्व

	Modifications Copyright 2020 Paulhindemith

	The original source code can be referenced from the link below.
	https://github.com/grafana-tools/sdk/blob/bdcab199ffdec390d845266c855ee01af90135a1/rest-request.go
	The change history can be obtained by looking at the differences from the
	following commit that added as the original source code.
	52e2c561d60ac579d97a5eabeaae42f0ce0db531
*/

package grafana_golang

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

// DefaultHTTPClient initialized Grafana with appropriate conditions.
// It allows you globally redefine HTTP client.
var DefaultHTTPClient = http.DefaultClient

// Client uses Grafana REST API for interacting with Grafana server.
type Client struct {
	baseURL string
	client  *http.Client
	reqOpts []func(req *http.Request)
}

func NewClientWithOpt(apiURL string, client *http.Client, reqOpts ...func(req *http.Request)) *Client {
	baseURL, _ := url.Parse(apiURL)
	return &Client{baseURL: baseURL.String(), client: client, reqOpts: reqOpts}
}

func (r *Client) Get(query string, params url.Values) ([]byte, int, error) {
	return r.DoRequest("GET", query, params, nil)
}

func (r *Client) Patch(query string, params url.Values, body []byte) ([]byte, int, error) {
	return r.DoRequest("PATCH", query, params, bytes.NewBuffer(body))
}

func (r *Client) Put(query string, params url.Values, body []byte) ([]byte, int, error) {
	return r.DoRequest("PUT", query, params, bytes.NewBuffer(body))
}

func (r *Client) Post(query string, params url.Values, body []byte) ([]byte, int, error) {
	return r.DoRequest("POST", query, params, bytes.NewBuffer(body))
}

func (r *Client) Delete(query string) ([]byte, int, error) {
	return r.DoRequest("DELETE", query, nil, nil)
}

func (r *Client) DoRequest(method, query string, params url.Values, buf io.Reader) ([]byte, int, error) {
	u, _ := url.Parse(r.baseURL)
	u.Path = path.Join(u.Path, query)
	if params != nil {
		u.RawQuery = params.Encode()
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, 0, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "autograf")

	for _, opt := range r.reqOpts {
		opt(req)
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return data, resp.StatusCode, err
}

func WithBasicAuth(user, password string) func(req *http.Request) {
	return func(req *http.Request) {
		req.SetBasicAuth(user, password)
	}
}

func WithBearerAuth(token string) func(req *http.Request) {
	return func(req *http.Request) {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	}
}
