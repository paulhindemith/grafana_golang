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
	https://github.com/grafana-tools/sdk/blob/bdcab199ffdec390d845266c855ee01af90135a1/panel.go
	The change history can be obtained by looking at the differences from the
	following commit that added as the original source code.
	52e2c561d60ac579d97a5eabeaae42f0ce0db531
*/

package main

type (
	// Panel represents panels of different types defined in Grafana.
	Panel struct {
		CommonPanel
		// Should be initialized only one type of panels.
		// OfType field defines which of types below will be used.
		*GraphPanel
		*TablePanel
		*TextPanel
		*SinglestatPanel
		*DashlistPanel
		*PluginlistPanel
		*RowPanel
		*AlertlistPanel
		*CustomPanel
	}
	panelType   int8
	CommonPanel struct {
		Datasource *string `json:"datasource,omitempty"` // metrics
		Editable   bool    `json:"editable"`
		Error      bool    `json:"error"`
		GridPos    struct {
			H *int `json:"h,omitempty"`
			W *int `json:"w,omitempty"`
			X *int `json:"x,omitempty"`
			Y *int `json:"y,omitempty"`
		} `json:"gridPos,omitempty"`
		Height           *string   `json:"height,omitempty"` // general
		HideTimeOverride *bool     `json:"hideTimeOverride,omitempty"`
		ID               uint      `json:"id"`
		IsNew            bool      `json:"isNew"`
		Links            []link    `json:"links,omitempty"`    // general
		MinSpan          *float32  `json:"minSpan,omitempty"`  // templating options
		OfType           panelType `json:"-"`                  // it required for defining type of the panel
		Renderer         *string   `json:"renderer,omitempty"` // display styles
		Repeat           *string   `json:"repeat,omitempty"`   // templating options
		// RepeatIteration *int64   `json:"repeatIteration,omitempty"`
		RepeatPanelID *uint `json:"repeatPanelId,omitempty"`
		ScopedVars    map[string]struct {
			Selected bool   `json:"selected"`
			Text     string `json:"text"`
			Value    string `json:"value"`
		} `json:"scopedVars,omitempty"`
		Span        float32 `json:"span"`  // general
		Title       string  `json:"title"` // general
		Transparent bool    `json:"transparent"`
		Type        string  `json:"type"`
		Alert       *Alert  `json:"alert,omitempty"`
	}
	AlertCondition struct {
		Evaluator struct {
			Params []float64 `json:"params,omitempty"`
			Type   string    `json:"type,omitempty"`
		} `json:"evaluator,omitempty"`
		Operator struct {
			Type string `json:"type,omitempty"`
		} `json:"operator,omitempty"`
		Query struct {
			Params []string `json:"params,omitempty"`
		} `json:"query,omitempty"`
		Reducer struct {
			Params []string `json:"params,omitempty"`
			Type   string   `json:"type,omitempty"`
		} `json:"reducer,omitempty"`
		Type string `json:"type,omitempty"`
	}
	Alert struct {
		Conditions          []AlertCondition    `json:"conditions,omitempty"`
		ExecutionErrorState string              `json:"executionErrorState,omitempty"`
		Frequency           string              `json:"frequency,omitempty"`
		Handler             int                 `json:"handler,omitempty"`
		Name                string              `json:"name,omitempty"`
		NoDataState         string              `json:"noDataState,omitempty"`
		Notifications       []AlertNotification `json:"notifications,omitempty"`
		Message             string              `json:"message,omitempty"`
		For                 string              `json:"for,omitempty"`
	}
	GraphPanel struct {
		AliasColors interface{} `json:"aliasColors"` // XXX
		Bars        bool        `json:"bars"`
		DashLength  *uint       `json:"dashLength,omitempty"`
		Dashes      *bool       `json:"dashes,omitempty"`
		Decimals    *uint       `json:"decimals,omitempty"`
		Description *string     `json:"description,omitempty"`
		Fill        int         `json:"fill"`
		//		Grid        grid        `json:"grid"` obsoleted in 4.1 by xaxis and yaxis

		Legend struct {
			AlignAsTable bool  `json:"alignAsTable"`
			Avg          bool  `json:"avg"`
			Current      bool  `json:"current"`
			HideEmpty    bool  `json:"hideEmpty"`
			HideZero     bool  `json:"hideZero"`
			Max          bool  `json:"max"`
			Min          bool  `json:"min"`
			RightSide    bool  `json:"rightSide"`
			Show         bool  `json:"show"`
			SideWidth    *uint `json:"sideWidth,omitempty"`
			Total        bool  `json:"total"`
			Values       bool  `json:"values"`
		} `json:"legend,omitempty"`
		LeftYAxisLabel  *string          `json:"leftYAxisLabel,omitempty"`
		Lines           bool             `json:"lines"`
		Linewidth       uint             `json:"linewidth"`
		NullPointMode   string           `json:"nullPointMode"`
		Percentage      bool             `json:"percentage"`
		Pointradius     int              `json:"pointradius"`
		Points          bool             `json:"points"`
		RightYAxisLabel *string          `json:"rightYAxisLabel,omitempty"`
		SeriesOverrides []SeriesOverride `json:"seriesOverrides,omitempty"`
		SpaceLength     *uint            `json:"spaceLength,omitempty"`
		Stack           bool             `json:"stack"`
		SteppedLine     bool             `json:"steppedLine"`
		Targets         []Target         `json:"targets,omitempty"`
		Thresholds      []Threshold      `json:"thresholds,omitempty"`
		TimeFrom        *string          `json:"timeFrom,omitempty"`
		TimeShift       *string          `json:"timeShift,omitempty"`
		Tooltip         Tooltip          `json:"tooltip"`
		XAxis           bool             `json:"x-axis,omitempty"`
		YAxis           bool             `json:"y-axis,omitempty"`
		YFormats        []string         `json:"y_formats,omitempty"`
		Xaxis           Axis             `json:"xaxis"` // was added in Grafana 4.x?
		Yaxes           []Axis           `json:"yaxes"` // was added in Grafana 4.x?
	}
	Threshold struct {
		// the alert threshold value, we do not omitempty, since 0 is a valid
		// threshold
		Value float32 `json:"value"`
		// critical, warning, ok, custom
		ColorMode string `json:"colorMode,omitempty"`
		// gt or lt
		Op   string `json:"op,omitempty"`
		Fill bool   `json:"fill"`
		Line bool   `json:"line"`
		// hexadecimal color (e.g. #629e51, only when ColorMode is "custom")
		FillColor string `json:"fillColor,omitempty"`
		// hexadecimal color (e.g. #629e51, only when ColorMode is "custom")
		LineColor string `json:"lineColor,omitempty"`
		// left or right
		Yaxis string `json:"yaxis,omitempty"`
	}

	Tooltip struct {
		Shared       bool   `json:"shared"`
		ValueType    string `json:"value_type"`
		MsResolution bool   `json:"msResolution,omitempty"` // was added in Grafana 3.x
		Sort         int    `json:"sort,omitempty"`
	}
	TablePanel struct {
		Columns []Column `json:"columns"`
		Sort    *struct {
			Col  uint `json:"col"`
			Desc bool `json:"desc"`
		} `json:"sort,omitempty"`
		Styles    []ColumnStyle `json:"styles"`
		Transform string        `json:"transform"`
		Targets   []Target      `json:"targets,omitempty"`
		Scroll    bool          `json:"scroll"` // from grafana 3.x
	}
	TextPanel struct {
		Content    string `json:"content"`
		Mode       string `json:"mode"`
		PageSize   uint   `json:"pageSize"`
		Scroll     bool   `json:"scroll"`
		ShowHeader bool   `json:"showHeader"`
		Sort       struct {
			Col  int  `json:"col"`
			Desc bool `json:"desc"`
		} `json:"sort"`
		Styles []ColumnStyle `json:"styles"`
	}
	SinglestatPanel struct {
		Colors          []string `json:"colors"`
		ColorValue      bool     `json:"colorValue"`
		ColorBackground bool     `json:"colorBackground"`
		Decimals        int      `json:"decimals"`
		Format          string   `json:"format"`
		Gauge           struct {
			MaxValue         float32 `json:"maxValue"`
			MinValue         float32 `json:"minValue"`
			Show             bool    `json:"show"`
			ThresholdLabels  bool    `json:"thresholdLabels"`
			ThresholdMarkers bool    `json:"thresholdMarkers"`
		} `json:"gauge,omitempty"`
		MappingType     *uint       `json:"mappingType,omitempty"`
		MappingTypes    []*MapType  `json:"mappingTypes,omitempty"`
		MaxDataPoints   *IntString  `json:"maxDataPoints,omitempty"`
		NullPointMode   string      `json:"nullPointMode"`
		Postfix         *string     `json:"postfix,omitempty"`
		PostfixFontSize *string     `json:"postfixFontSize,omitempty"`
		Prefix          *string     `json:"prefix,omitempty"`
		PrefixFontSize  *string     `json:"prefixFontSize,omitempty"`
		RangeMaps       []*RangeMap `json:"rangeMaps,omitempty"`
		SparkLine       struct {
			FillColor *string  `json:"fillColor,omitempty"`
			Full      bool     `json:"full,omitempty"`
			LineColor *string  `json:"lineColor,omitempty"`
			Show      bool     `json:"show,omitempty"`
			YMin      *float64 `json:"ymin,omitempty"`
			YMax      *float64 `json:"ymax,omitempty"`
		} `json:"sparkline,omitempty"`
		Targets       []Target   `json:"targets,omitempty"`
		Thresholds    string     `json:"thresholds"`
		ValueFontSize string     `json:"valueFontSize"`
		ValueMaps     []ValueMap `json:"valueMaps"`
		ValueName     string     `json:"valueName"`
	}
	DashlistPanel struct {
		Mode  string   `json:"mode"`
		Limit uint     `json:"limit"`
		Query string   `json:"query"`
		Tags  []string `json:"tags"`
	}
	PluginlistPanel struct {
		Limit int `json:"limit,omitempty"`
	}
	AlertlistPanel struct {
		OnlyAlertsOnDashboard bool     `json:"onlyAlertsOnDashboard"`
		Show                  string   `json:"show"`
		SortOrder             int      `json:"sortOrder"`
		Limit                 int      `json:"limit"`
		StateFilter           []string `json:"stateFilter"`
		NameFilter            string   `json:"nameFilter,omitempty"`
		DashboardTags         []string `json:"dashboardTags,omitempty"`
	}
	RowPanel struct {
		Panels []Panel
	}
	CustomPanel map[string]interface{}
)

// for a graph panel
type (
	// TODO look at schema versions carefully
	// grid was obsoleted by xaxis and yaxes
	grid struct {
		LeftLogBase     *int     `json:"leftLogBase"`
		LeftMax         *int     `json:"leftMax"`
		LeftMin         *int     `json:"leftMin"`
		RightLogBase    *int     `json:"rightLogBase"`
		RightMax        *int     `json:"rightMax"`
		RightMin        *int     `json:"rightMin"`
		Threshold1      *float64 `json:"threshold1"`
		Threshold1Color string   `json:"threshold1Color"`
		Threshold2      *float64 `json:"threshold2"`
		Threshold2Color string   `json:"threshold2Color"`
		ThresholdLine   bool     `json:"thresholdLine"`
	}
	xaxis struct {
		Mode   string      `json:"mode"`
		Name   interface{} `json:"name"` // TODO what is this?
		Show   bool        `json:"show"`
		Values *[]string   `json:"values,omitempty"`
	}
	Axis struct {
		Format  string       `json:"format"`
		LogBase int          `json:"logBase"`
		Max     *FloatString `json:"max,omitempty"`
		Min     *FloatString `json:"min,omitempty"`
		Show    bool         `json:"show"`
		Label   string       `json:"label,omitempty"`
	}
	SeriesOverride struct {
		Alias         string      `json:"alias"`
		Bars          *bool       `json:"bars,omitempty"`
		Color         *string     `json:"color,omitempty"`
		Fill          *int        `json:"fill,omitempty"`
		FillBelowTo   *string     `json:"fillBelowTo,omitempty"`
		Legend        *bool       `json:"legend,omitempty"`
		Lines         *bool       `json:"lines,omitempty"`
		Stack         *BoolString `json:"stack,omitempty"`
		Transform     *string     `json:"transform,omitempty"`
		YAxis         *int        `json:"yaxis,omitempty"`
		ZIndex        *int        `json:"zindex,omitempty"`
		NullPointMode *string     `json:"nullPointMode,omitempty"`
	}
)

// for a table
type (
	Column struct {
		TextType string `json:"text"`
		Value    string `json:"value"`
	}
	ColumnStyle struct {
		Alias      *string   `json:"alias"`
		DateFormat *string   `json:"dateFormat,omitempty"`
		Pattern    string    `json:"pattern"`
		Type       string    `json:"type"`
		ColorMode  *string   `json:"colorMode,omitempty"`
		Colors     *[]string `json:"colors,omitempty"`
		Decimals   *uint     `json:"decimals,omitempty"`
		Thresholds *[]string `json:"thresholds,omitempty"`
		Unit       *string   `json:"unit,omitempty"`
	}
)

// for a singlestat
type ValueMap struct {
	Op       string `json:"op"`
	TextType string `json:"text"`
	Value    string `json:"value"`
}

// for an any panel
type Target struct {
	RefID      string `json:"refId"`
	Datasource string `json:"datasource,omitempty"`

	// For Prometheus
	Expr           string `json:"expr,omitempty"`
	IntervalFactor int    `json:"intervalFactor,omitempty"`
	Interval       string `json:"interval,omitempty"`
	Step           int    `json:"step,omitempty"`
	LegendFormat   string `json:"legendFormat,omitempty"`
	Instant        bool   `json:"instant,omitempty"`
	Format         string `json:"format,omitempty"`

	// For Elasticsearch
	DsType  *string `json:"dsType,omitempty"`
	Metrics []struct {
		ID    string `json:"id"`
		Field string `json:"field"`
		Type  string `json:"type"`
	} `json:"metrics,omitempty"`
	Query      string `json:"query,omitempty"`
	Alias      string `json:"alias,omitempty"`
	RawQuery   bool   `json:"rawQuery,omitempty"`
	TimeField  string `json:"timeField,omitempty"`
	BucketAggs []struct {
		ID       string `json:"id"`
		Field    string `json:"field"`
		Type     string `json:"type"`
		Settings struct {
			Interval    string `json:"interval,omitempty"`
			MinDocCount int    `json:"min_doc_count"`
			Order       string `json:"order,omitempty"`
			OrderBy     string `json:"orderBy,omitempty"`
			Size        string `json:"size,omitempty"`
		} `json:"settings"`
	} `json:"bucketAggs,omitempty"`

	// For Graphite
	Target string `json:"target,omitempty"`

	// For CloudWatch
	Namespace  string            `json:"namespace,omitempty"`
	MetricName string            `json:"metricName,omitempty"`
	Statistics []string          `json:"statistics,omitempty"`
	Dimensions map[string]string `json:"dimensions,omitempty"`
	Period     string            `json:"period,omitempty"`
	Region     string            `json:"region,omitempty"`
}

type MapType struct {
	Name  *string `json:"name,omitempty"`
	Value *int    `json:"value,omitempty"`
}

type RangeMap struct {
	From *string `json:"from,omitempty"`
	Text *string `json:"text,omitempty"`
	To   *string `json:"to,omitempty"`
}
