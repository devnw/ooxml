package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_ScatterStyleConstructor(t *testing.T) {
	v := chart.NewCT_ScatterStyle()
	if v == nil {
		t.Errorf("chart.NewCT_ScatterStyle must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_ScatterStyle should validate: %s", err)
	}
}

func TestCT_ScatterStyleMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_ScatterStyle()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_ScatterStyle()
	xml.Unmarshal(buf, v2)
}
