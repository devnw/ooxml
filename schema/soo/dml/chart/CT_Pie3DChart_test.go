package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_Pie3DChartConstructor(t *testing.T) {
	v := chart.NewCT_Pie3DChart()
	if v == nil {
		t.Errorf("chart.NewCT_Pie3DChart must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_Pie3DChart should validate: %s", err)
	}
}

func TestCT_Pie3DChartMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_Pie3DChart()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_Pie3DChart()
	xml.Unmarshal(buf, v2)
}
