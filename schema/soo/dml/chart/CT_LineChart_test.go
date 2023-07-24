package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_LineChartConstructor(t *testing.T) {
	v := chart.NewCT_LineChart()
	if v == nil {
		t.Errorf("chart.NewCT_LineChart must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_LineChart should validate: %s", err)
	}
}

func TestCT_LineChartMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_LineChart()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_LineChart()
	xml.Unmarshal(buf, v2)
}
