package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_ChartConstructor(t *testing.T) {
	v := chart.NewCT_Chart()
	if v == nil {
		t.Errorf("chart.NewCT_Chart must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_Chart should validate: %s", err)
	}
}

func TestCT_ChartMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_Chart()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_Chart()
	xml.Unmarshal(buf, v2)
}
