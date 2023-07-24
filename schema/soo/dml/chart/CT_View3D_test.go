package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_View3DConstructor(t *testing.T) {
	v := chart.NewCT_View3D()
	if v == nil {
		t.Errorf("chart.NewCT_View3D must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_View3D should validate: %s", err)
	}
}

func TestCT_View3DMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_View3D()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_View3D()
	xml.Unmarshal(buf, v2)
}
