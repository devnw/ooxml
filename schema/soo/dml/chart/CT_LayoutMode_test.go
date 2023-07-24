package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_LayoutModeConstructor(t *testing.T) {
	v := chart.NewCT_LayoutMode()
	if v == nil {
		t.Errorf("chart.NewCT_LayoutMode must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_LayoutMode should validate: %s", err)
	}
}

func TestCT_LayoutModeMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_LayoutMode()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_LayoutMode()
	xml.Unmarshal(buf, v2)
}
