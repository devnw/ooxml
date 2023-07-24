package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_OfPieChartConstructor(t *testing.T) {
	v := chart.NewCT_OfPieChart()
	if v == nil {
		t.Errorf("chart.NewCT_OfPieChart must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_OfPieChart should validate: %s", err)
	}
}

func TestCT_OfPieChartMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_OfPieChart()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_OfPieChart()
	xml.Unmarshal(buf, v2)
}
