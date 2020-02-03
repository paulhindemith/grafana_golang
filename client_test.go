/*
Copyright 2020 Paulhindemith

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

package grafana_golang

import (
	"net/http"
	"testing"
)

func TestClientWithBasicAuth(t *testing.T) {
	var testcases = []struct {
		reqOpt        func(req *http.Request)
		expectedError bool
		errorMsg      string
	}{
		{
			reqOpt:        func(req *http.Request) {},
			expectedError: true,
			errorMsg:      "HTTP error 425: returns unauthorizeds",
		},
		{
			reqOpt:        WithBasicAuth("admin", "admin"),
			expectedError: false,
		},
	}

	for _, tc := range testcases {
		client := NewClientWithOpt(TestEndpoint, http.DefaultClient, tc.reqOpt)
		apiKeyInfo := &APIKeyInfo{
			Name:          "testKey",
			Role:          AdminRole,
			SecondsToLive: 0,
		}
		apiKey, err := client.CreateAPIKey(apiKeyInfo)
		if tc.expectedError && err == nil {
			t.Fatal("error does not exist.")
		} else if !tc.expectedError && err != nil {
			t.Fatal(err.Error())
		}
		if apiKey != nil {
			if err := TeardownAPIKey(client); err != nil {
				t.Fatal(err.Error())
			}
		}
	}
}

func TestClientWithBearerAuth(t *testing.T) {
	var apiKeyInfo *APIKeyInfo

	client := NewClientWithOpt(TestEndpoint, http.DefaultClient, WithBasicAuth("admin", "admin"))
	apiKeyInfo = &APIKeyInfo{
		Name:          "testKey",
		Role:          AdminRole,
		SecondsToLive: 0,
	}
	apiKey, err := client.CreateAPIKey(apiKeyInfo)
	if err != nil {
		t.Fatal(err.Error())
	}

	// then
	apiKeyInfo = &APIKeyInfo{
		Name:          "testKeyAlt",
		Role:          AdminRole,
		SecondsToLive: 0,
	}
	interestClient := NewClientWithOpt(TestEndpoint, http.DefaultClient, WithBearerAuth(apiKey.Key))
	_, err = interestClient.CreateAPIKey(apiKeyInfo)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer func() {
		if err := TeardownAPIKey(client); err != nil {
			t.Fatal(err.Error())
		}
	}()
}
