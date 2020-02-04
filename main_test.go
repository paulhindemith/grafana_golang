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
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
)

const (
	grafanaVersion = "6.6.0"
	TestEndpoint   = "http://localhost:3000"
)

func TestMain(m *testing.M) {
	setup()
	ret := m.Run()
	teardown()

	os.Exit(ret)
}

func setup() {
	client := NewClientWithOpt(TestEndpoint, http.DefaultClient)
	if perr := wait.PollImmediate(time.Second, 30*time.Second, func() (bool, error) {
		res, err := client.Health()
		if err != nil {
			return false, err
		}

		switch {
		case res.Database != DatabaseStatusOK:
			return false, nil
		case res.Version != grafanaVersion:
			return false, fmt.Errorf("Grafana version is wrong. Expect: %s, Got: %s", grafanaVersion, res.Version)
		default:
			return true, nil
		}
	}); perr != nil {
		fmt.Print(perr.Error())
		return
	}
	if err := Teardown(client); err != nil {
		fmt.Print(err.Error())
		return
	}
}

func teardown() {
	client := NewClientWithOpt(TestEndpoint, http.DefaultClient, WithBasicAuth("admin", "admin"))

	if err := Teardown(client); err != nil {
		fmt.Print(err.Error())
		return
	}
}
