package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_MarkerStyleConstructor(t *testing.T) {
	v := chart.NewCT_MarkerStyle()
	if v == nil {
		t.Errorf("chart.NewCT_MarkerStyle must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_MarkerStyle should validate: %s", err)
	}
}

func TestCT_MarkerStyleMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_MarkerStyle()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_MarkerStyle()
	xml.Unmarshal(buf, v2)
}
