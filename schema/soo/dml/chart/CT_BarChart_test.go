package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_BarChartConstructor(t *testing.T) {
	v := chart.NewCT_BarChart()
	if v == nil {
		t.Errorf("chart.NewCT_BarChart must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_BarChart should validate: %s", err)
	}
}

func TestCT_BarChartMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_BarChart()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_BarChart()
	xml.Unmarshal(buf, v2)
}
