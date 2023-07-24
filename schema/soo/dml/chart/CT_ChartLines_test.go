package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_ChartLinesConstructor(t *testing.T) {
	v := chart.NewCT_ChartLines()
	if v == nil {
		t.Errorf("chart.NewCT_ChartLines must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_ChartLines should validate: %s", err)
	}
}

func TestCT_ChartLinesMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_ChartLines()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_ChartLines()
	xml.Unmarshal(buf, v2)
}
