package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_RotXConstructor(t *testing.T) {
	v := chart.NewCT_RotX()
	if v == nil {
		t.Errorf("chart.NewCT_RotX must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_RotX should validate: %s", err)
	}
}

func TestCT_RotXMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_RotX()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_RotX()
	xml.Unmarshal(buf, v2)
}
