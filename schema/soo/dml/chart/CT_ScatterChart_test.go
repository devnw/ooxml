package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_ScatterChartConstructor(t *testing.T) {
	v := chart.NewCT_ScatterChart()
	if v == nil {
		t.Errorf("chart.NewCT_ScatterChart must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_ScatterChart should validate: %s", err)
	}
}

func TestCT_ScatterChartMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_ScatterChart()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_ScatterChart()
	xml.Unmarshal(buf, v2)
}
