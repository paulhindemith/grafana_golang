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
	"io/ioutil"
)

func ReadTextFile(file string) (string, error) {
	var (
		err error
		raw []byte
	)
	if raw, err = ioutil.ReadFile(file); err != nil {
		return "", err
	}
	return string(raw), nil
}

func ReadDashboardFile(file string) (*Dashboard, error) {
	var (
		err       error
		raw       []byte
		dashboard Dashboard
	)
	if raw, err = ioutil.ReadFile(file); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(raw, &dashboard); err != nil {
		return nil, err
	}
	return &dashboard, nil
}

func Teardown(client *Client) error {
	var err error
	if err = TeardownAPIKey(client); err != nil {
		return err
	}
	if err = TeardownSnapshot(client); err != nil {
		return err
	}
	if err = TeardownDatasource(client); err != nil {
		return err
	}
	return nil
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

func TeardownDatasource(client *Client) error {
	datasoruces, err := client.GetDatasources()
	if err != nil {
		return err
	}
	for _, d := range datasoruces {
		if err := client.DeleteDatasource(d.ID); err != nil {
			return err
		}
	}
	return nil
}
