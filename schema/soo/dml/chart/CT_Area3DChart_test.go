package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_Area3DChartConstructor(t *testing.T) {
	v := chart.NewCT_Area3DChart()
	if v == nil {
		t.Errorf("chart.NewCT_Area3DChart must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_Area3DChart should validate: %s", err)
	}
}

func TestCT_Area3DChartMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_Area3DChart()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_Area3DChart()
	xml.Unmarshal(buf, v2)
}
