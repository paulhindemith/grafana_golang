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

func TestCreateAPIKey(t *testing.T) {
	client := NewClientWithOpt(TestEndpoint, http.DefaultClient, WithBasicAuth("admin", "admin"))
	apiKeyInfo := &APIKeyInfo{
		Name:          "testKeyAlt",
		Role:          AdminRole,
		SecondsToLive: 0,
	}
	_, err := client.CreateAPIKey(apiKeyInfo)
	if err != nil {
		t.Fatal(err.Error())
	}

	defer func() {
		if err := TeardownAPIKey(client); err != nil {
			t.Fatal(err.Error())
		}
	}()
}

func TeardownAPIKey(client *Client) error {
	apiKeyInfos, err := client.GetAPIKeys()
	if err != nil {
		return err
	}
	for _, i := range apiKeyInfos {
		if err := client.DeleteAPIKey(i.ID); err != nil {
			return err
		}
	}
	return nil
}
