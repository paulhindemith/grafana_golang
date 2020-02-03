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
	"time"
)

type Snapshot struct {
	Dashboard *Dashboard `json:"dashboard"`
	Name      string     `json:"name"`
	Expires   int        `json:"expires"`
	External  bool       `json:"external"`
}

type SnapshotInfo struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Key         string    `json:"key"`
	OrgId       uint      `json:"orgId"`
	UserId      uint      `json:"userId"`
	External    bool      `json:"external"`
	ExternalUrl string    `json:"externalUrl"`
	Expires     time.Time `json:"expires"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
}

type SnapshotResult struct {
	DeleteKey string `json:"deleteKey"`
	DeleteUrl string `json:"deleteUrl"`
	Key       string `json:"key"`
	Url       string `json:"url"`
}

func (r *Client) GetSnapshots() ([]*SnapshotInfo, error) {
	var (
		raw  []byte
		resp []*SnapshotInfo
		code int
		err  error
	)
	if raw, code, err = r.Get("api/dashboard/snapshots", nil); err != nil {
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

func (r *Client) CreateSnapshot(s *Snapshot) (*SnapshotResult, error) {
	var (
		raw  []byte
		resp SnapshotResult
		code int
		err  error
	)
	if raw, err = json.Marshal(s); err != nil {
		return nil, err
	}

	if raw, code, err = r.Post("api/snapshots", nil, raw); err != nil {
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

func (r *Client) DeleteSnapshot(key string) error {
	var (
		raw  []byte
		code int
		err  error
	)
	if raw, code, err = r.Delete(fmt.Sprintf("api/snapshots/%s", key)); err != nil {
		return err
	}

	if code != http.StatusOK {
		return fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}

	return nil
}
