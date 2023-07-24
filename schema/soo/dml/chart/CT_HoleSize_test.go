package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_HoleSizeConstructor(t *testing.T) {
	v := chart.NewCT_HoleSize()
	if v == nil {
		t.Errorf("chart.NewCT_HoleSize must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_HoleSize should validate: %s", err)
	}
}

func TestCT_HoleSizeMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_HoleSize()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_HoleSize()
	xml.Unmarshal(buf, v2)
}
