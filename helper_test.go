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
	"os"
	"path"
	"testing"
)

func TestReadTokenFile(t *testing.T) {
	gp := os.Getenv("GOPATH")
	tp := path.Join(gp, "src/github.com/paulhindemith/grafana-client/test/testdata-token.txt")
	token, err := ReadTokenFile(tp)
	if err != nil {
		t.Fatal(err.Error())
	}
	if token == "" {
		t.Fatal("Token is not read well.")
	}
}

func TestReadDashboardFile(t *testing.T) {
	gp := os.Getenv("GOPATH")
	dp := path.Join(gp, "src/github.com/paulhindemith/grafana-client/test/testdata-dashboard.json")
	db, err := ReadDashboardFile(dp)
	if err != nil {
		t.Fatal(err.Error())
	}
	if db.UID == "" {
		t.Fatal("Dashboard is not read well.")
	}
}
