package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_SerAxConstructor(t *testing.T) {
	v := chart.NewCT_SerAx()
	if v == nil {
		t.Errorf("chart.NewCT_SerAx must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_SerAx should validate: %s", err)
	}
}

func TestCT_SerAxMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_SerAx()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_SerAx()
	xml.Unmarshal(buf, v2)
}
