package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_AreaChartConstructor(t *testing.T) {
	v := chart.NewCT_AreaChart()
	if v == nil {
		t.Errorf("chart.NewCT_AreaChart must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_AreaChart should validate: %s", err)
	}
}

func TestCT_AreaChartMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_AreaChart()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_AreaChart()
	xml.Unmarshal(buf, v2)
}
