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
	"encoding/json"
	"fmt"
	"net/http"
)

type HealthResult struct {
	Commit   string         `json:"commit"`
	Database DatabaseStatus `json:"database"`
	Version  string         `json:"version"`
}

type DatabaseStatus string

const (
	DatabaseStatusOK DatabaseStatus = "ok"
)

func (r *Client) Health() (*HealthResult, error) {
	var err error

	raw, code, err := r.Get("api/health", nil)
	if err != nil {
		return nil, err
	}

	if code != http.StatusOK {
		return nil, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}

	var resp HealthResult
	if err = json.Unmarshal(raw, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
