package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_PerspectiveConstructor(t *testing.T) {
	v := chart.NewCT_Perspective()
	if v == nil {
		t.Errorf("chart.NewCT_Perspective must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_Perspective should validate: %s", err)
	}
}

func TestCT_PerspectiveMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_Perspective()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_Perspective()
	xml.Unmarshal(buf, v2)
}
