package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_ShapeNonVisualConstructor(t *testing.T) {
	v := pml.NewCT_ShapeNonVisual()
	if v == nil {
		t.Errorf("pml.NewCT_ShapeNonVisual must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_ShapeNonVisual should validate: %s", err)
	}
}

func TestCT_ShapeNonVisualMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_ShapeNonVisual()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_ShapeNonVisual()
	xml.Unmarshal(buf, v2)
}
