package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestChartSpaceConstructor(t *testing.T) {
	v := chart.NewChartSpace()
	if v == nil {
		t.Errorf("chart.NewChartSpace must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.ChartSpace should validate: %s", err)
	}
}

func TestChartSpaceMarshalUnmarshal(t *testing.T) {
	v := chart.NewChartSpace()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewChartSpace()
	xml.Unmarshal(buf, v2)
}
