package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_ShapeConstructor(t *testing.T) {
	v := chart.NewCT_Shape()
	if v == nil {
		t.Errorf("chart.NewCT_Shape must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_Shape should validate: %s", err)
	}
}

func TestCT_ShapeMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_Shape()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_Shape()
	xml.Unmarshal(buf, v2)
}
