package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_RadarChartConstructor(t *testing.T) {
	v := chart.NewCT_RadarChart()
	if v == nil {
		t.Errorf("chart.NewCT_RadarChart must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_RadarChart should validate: %s", err)
	}
}

func TestCT_RadarChartMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_RadarChart()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_RadarChart()
	xml.Unmarshal(buf, v2)
}
