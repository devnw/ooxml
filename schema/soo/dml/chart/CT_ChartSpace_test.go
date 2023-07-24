package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_ChartSpaceConstructor(t *testing.T) {
	v := chart.NewCT_ChartSpace()
	if v == nil {
		t.Errorf("chart.NewCT_ChartSpace must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_ChartSpace should validate: %s", err)
	}
}

func TestCT_ChartSpaceMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_ChartSpace()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_ChartSpace()
	xml.Unmarshal(buf, v2)
}
