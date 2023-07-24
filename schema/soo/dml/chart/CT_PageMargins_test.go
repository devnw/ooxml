package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_PageMarginsConstructor(t *testing.T) {
	v := chart.NewCT_PageMargins()
	if v == nil {
		t.Errorf("chart.NewCT_PageMargins must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_PageMargins should validate: %s", err)
	}
}

func TestCT_PageMarginsMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_PageMargins()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_PageMargins()
	xml.Unmarshal(buf, v2)
}
