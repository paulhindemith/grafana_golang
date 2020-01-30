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

package main

import (
	"net/http"
	"os"
	"path"
	"testing"
)

func TestCreateSnapshot(t *testing.T) {
	client := NewClientWithOpt(TestEndpoint, http.DefaultClient, WithBasicAuth("admin", "admin"))
	gp := os.Getenv("GOPATH")
	dp := path.Join(gp, "src/github.com/paulhindemith/grafana-client/test/testdata-dashboard.json")
	dashboard, err := ReadDashboardFile(dp)
	if err != nil {
		t.Fatal(err.Error())
	}
	snapshot := &Snapshot{
		Dashboard: dashboard,
		Name:      "testSnapshot",
		Expires:   0,
		External:  false,
	}
	_, err = client.CreateSnapshot(snapshot)
	if err != nil {
		t.Fatal(err.Error())
	}

	defer func() {
		if err := TeardownSnapshot(client); err != nil {
			t.Fatal(err.Error())
		}
	}()
}

func TeardownSnapshot(client *Client) error {
	snapshots, err := client.GetSnapshots()
	if err != nil {
		return err
	}
	for _, s := range snapshots {
		if err := client.DeleteSnapshot(s.Key); err != nil {
			return err
		}
	}
	return nil
}
