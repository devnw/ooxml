package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_ManualLayoutConstructor(t *testing.T) {
	v := chart.NewCT_ManualLayout()
	if v == nil {
		t.Errorf("chart.NewCT_ManualLayout must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_ManualLayout should validate: %s", err)
	}
}

func TestCT_ManualLayoutMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_ManualLayout()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_ManualLayout()
	xml.Unmarshal(buf, v2)
}
