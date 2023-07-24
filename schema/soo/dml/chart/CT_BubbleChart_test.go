package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_BubbleChartConstructor(t *testing.T) {
	v := chart.NewCT_BubbleChart()
	if v == nil {
		t.Errorf("chart.NewCT_BubbleChart must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_BubbleChart should validate: %s", err)
	}
}

func TestCT_BubbleChartMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_BubbleChart()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_BubbleChart()
	xml.Unmarshal(buf, v2)
}
