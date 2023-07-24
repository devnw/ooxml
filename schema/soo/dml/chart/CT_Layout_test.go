package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_LayoutConstructor(t *testing.T) {
	v := chart.NewCT_Layout()
	if v == nil {
		t.Errorf("chart.NewCT_Layout must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_Layout should validate: %s", err)
	}
}

func TestCT_LayoutMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_Layout()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_Layout()
	xml.Unmarshal(buf, v2)
}
