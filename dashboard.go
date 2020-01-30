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
	https://github.com/grafana-tools/sdk/blob/bdcab199ffdec390d845266c855ee01af90135a1/rest-dashboard.go
	The change history can be obtained by looking at the differences from the
	following commit that added as the original source code.
	52e2c561d60ac579d97a5eabeaae42f0ce0db531
*/

package main

type (
	// Board represents Grafana dashboard.
	Dashboard struct {
		ID              uint       `json:"id,omitempty"`
		UID             string     `json:"uid,omitempty"`
		Slug            string     `json:"slug"`
		Title           string     `json:"title"`
		OriginalTitle   string     `json:"originalTitle"`
		Tags            []string   `json:"tags"`
		Style           string     `json:"style"`
		Timezone        string     `json:"timezone"`
		Editable        bool       `json:"editable"`
		HideControls    bool       `json:"hideControls" graf:"hide-controls"`
		SharedCrosshair bool       `json:"sharedCrosshair" graf:"shared-crosshair"`
		Panels          []*Panel   `json:"panels"`
		Rows            []*Row     `json:"rows"`
		Templating      Templating `json:"templating"`
		Annotations     struct {
			List []Annotation `json:"list"`
		} `json:"annotations"`
		Refresh       *BoolString `json:"refresh,omitempty"`
		SchemaVersion uint        `json:"schemaVersion"`
		Version       uint        `json:"version"`
		Links         []link      `json:"links"`
		Time          Time        `json:"time"`
		Timepicker    Timepicker  `json:"timepicker"`
		lastPanelID   uint
		GraphTooltip  int `json:"graphTooltip,omitempty"`
	}
	Time struct {
		From string `json:"from"`
		To   string `json:"to"`
	}
	Timepicker struct {
		Now              *bool    `json:"now,omitempty"`
		RefreshIntervals []string `json:"refresh_intervals"`
		TimeOptions      []string `json:"time_options"`
	}
	Templating struct {
		List []TemplateVar `json:"list"`
	}
	TemplateVar struct {
		Name        string   `json:"name"`
		Type        string   `json:"type"`
		Auto        bool     `json:"auto,omitempty"`
		AutoCount   *int     `json:"auto_count,omitempty"`
		Datasource  *string  `json:"datasource"`
		Refresh     BoolInt  `json:"refresh"`
		Options     []Option `json:"options"`
		IncludeAll  bool     `json:"includeAll"`
		AllFormat   string   `json:"allFormat"`
		AllValue    string   `json:"allValue"`
		Multi       bool     `json:"multi"`
		MultiFormat string   `json:"multiFormat"`
		Query       string   `json:"query"`
		Regex       string   `json:"regex"`
		Current     Current  `json:"current"`
		Label       string   `json:"label"`
		Hide        uint8    `json:"hide"`
		Sort        int      `json:"sort"`
	}
	// for templateVar
	Option struct {
		Text     string `json:"text"`
		Value    string `json:"value"`
		Selected bool   `json:"selected"`
	}
	// for templateVar
	Current struct {
		Tags  []*string   `json:"tags,omitempty"`
		Text  string      `json:"text"`
		Value interface{} `json:"value"` // TODO select more precise type
	}
	Annotation struct {
		Name       string   `json:"name"`
		Datasource *string  `json:"datasource"`
		ShowLine   bool     `json:"showLine"`
		IconColor  string   `json:"iconColor"`
		LineColor  string   `json:"lineColor"`
		IconSize   uint     `json:"iconSize"`
		Enable     bool     `json:"enable"`
		Query      string   `json:"query"`
		TextField  string   `json:"textField"`
		TagsField  string   `json:"tagsField"`
		Tags       []string `json:"tags"`
		Type       string   `json:"type"`
	}
)

// link represents link to another dashboard or external weblink
type link struct {
	Title       string   `json:"title"`
	Type        string   `json:"type"`
	AsDropdown  *bool    `json:"asDropdown,omitempty"`
	DashURI     *string  `json:"dashUri,omitempty"`
	Dashboard   *string  `json:"dashboard,omitempty"`
	Icon        *string  `json:"icon,omitempty"`
	IncludeVars bool     `json:"includeVars"`
	KeepTime    *bool    `json:"keepTime,omitempty"`
	Params      *string  `json:"params,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	TargetBlank *bool    `json:"targetBlank,omitempty"`
	Tooltip     *string  `json:"tooltip,omitempty"`
	URL         *string  `json:"url,omitempty"`
}

// Height of rows maybe passed as number (ex 200) or
// as string (ex "200px") or empty string
type Height string
