package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_DoughnutChartConstructor(t *testing.T) {
	v := chart.NewCT_DoughnutChart()
	if v == nil {
		t.Errorf("chart.NewCT_DoughnutChart must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_DoughnutChart should validate: %s", err)
	}
}

func TestCT_DoughnutChartMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_DoughnutChart()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_DoughnutChart()
	xml.Unmarshal(buf, v2)
}
