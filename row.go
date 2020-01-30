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
	https://github.com/grafana-tools/sdk/blob/bdcab199ffdec390d845266c855ee01af90135a1/row.go
	The change history can be obtained by looking at the differences from the
	following commit that added as the original source code.
	52e2c561d60ac579d97a5eabeaae42f0ce0db531
*/

package main

// Row represents single row of Grafana dashboard.
type Row struct {
	Title     string  `json:"title"`
	ShowTitle bool    `json:"showTitle"`
	Collapse  bool    `json:"collapse"`
	Editable  bool    `json:"editable"`
	Height    Height  `json:"height"`
	Panels    []Panel `json:"panels"`
	Repeat    *string `json:"repeat"`
}
