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
	https://github.com/grafana-tools/sdk/blob/bdcab199ffdec390d845266c855ee01af90135a1/rest-datasource.go
	The change history can be obtained by looking at the differences from the
	following commit that added as the original source code.
	52e2c561d60ac579d97a5eabeaae42f0ce0db531
*/

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Datasource as described in the doc
// http://docs.grafana.org/reference/http_api/#get-all-datasources
type Datasource struct {
	ID                uint        `json:"id"`
	OrgID             uint        `json:"orgId"`
	Name              string      `json:"name"`
	Type              string      `json:"type"`
	TypeLogoURL       string      `json:"typeLogoUrl"`
	Access            string      `json:"access"` // direct or proxy
	URL               string      `json:"url"`
	Password          string      `json:"password,omitempty"`
	User              string      `json:"user,omitempty"`
	Database          string      `json:"database,omitempty"`
	BasicAuth         bool        `json:"basicAuth,omitempty"`
	BasicAuthUser     string      `json:"basicAuthUser,omitempty"`
	BasicAuthPassword string      `json:"basicAuthPassword,omitempty"`
	IsDefault         bool        `json:"isDefault"`
	JSONData          interface{} `json:"jsonData"`
	SecureJSONData    interface{} `json:"secureJsonData"`
	Version           interface{} `json:"version,omitempty"`
	ReadOnly          interface{} `json:"readOnly,omitempty"`
}

type DatasourceResult struct {
	ID         uint        `json:"id"`
	Name       string      `json:"name"`
	Message    string      `json:"message"`
	Datasource *Datasource `json:"datasource"`
}

func (r *Client) GetDatasources() ([]*Datasource, error) {
	var (
		raw  []byte
		resp []*Datasource
		code int
		err  error
	)
	if raw, code, err = r.Get("api/datasources", nil); err != nil {
		return nil, err
	}

	if code != http.StatusOK {
		return nil, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}

	if err = json.Unmarshal(raw, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateDatasource creates a new datasource.
// Reflects POST /api/datasources API call.
func (r *Client) CreateDatasource(ds *Datasource) (*DatasourceResult, error) {
	var (
		raw  []byte
		resp DatasourceResult
		code int
		err  error
	)
	if raw, err = json.Marshal(ds); err != nil {
		return nil, err
	}
	if raw, code, err = r.Post("api/datasources", nil, raw); err != nil {
		return nil, err
	}

	if code != http.StatusOK {
		return nil, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}

	if err = json.Unmarshal(raw, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (r *Client) DeleteDatasource(id uint) error {
	var (
		raw  []byte
		code int
		err  error
	)
	if raw, code, err = r.Delete(fmt.Sprintf("api/datasources/%d", id)); err != nil {
		return err
	}

	if code != http.StatusOK {
		return fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}

	return nil
}
