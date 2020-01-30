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
	"encoding/json"
	"fmt"
	"net/http"
)

type APIKeyInfo struct {
	ID            int        `json:"id, omitempty"`
	Name          string     `json:"name"`
	Role          APIKeyRole `json:"role"`
	SecondsToLive int        `json:"secondsToLive"`
}

type APIKeyRole string

const (
	AdminRole  APIKeyRole = "Admin"
	ViewerRole APIKeyRole = "Viewer"
	EditorRole APIKeyRole = "Editor"
)

type APIKey struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

func (r *Client) GetAPIKeys() ([]*APIKeyInfo, error) {
	var (
		raw  []byte
		code int
		err  error
		resp []*APIKeyInfo
	)
	if raw, code, err = r.Get("api/auth/keys", nil); err != nil {
		return nil, err
	}
	if code != http.StatusOK {
		return nil, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}
	if err = json.Unmarshal(raw, &resp); err != nil {
		return nil, err
	}
	return resp, err
}

func (r *Client) CreateAPIKey(req *APIKeyInfo) (*APIKey, error) {
	var (
		raw    []byte
		apiKey APIKey
		code   int
		err    error
	)
	if raw, err = json.Marshal(req); err != nil {
		return nil, err
	}
	if raw, code, err = r.Post("api/auth/keys", nil, raw); err != nil {
		return nil, err
	}
	if code != http.StatusOK {
		return nil, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}
	if err = json.Unmarshal(raw, &apiKey); err != nil {
		return nil, err
	}
	return &apiKey, err
}

func (r *Client) DeleteAPIKey(id int) error {
	var (
		raw  []byte
		code int
		err  error
	)
	if _, code, err = r.Delete(fmt.Sprintf("api/auth/keys/%d", id)); err != nil {
		return err
	}
	if code != http.StatusOK {
		return fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}
	return nil
}
