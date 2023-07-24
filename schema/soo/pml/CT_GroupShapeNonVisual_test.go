package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_GroupShapeNonVisualConstructor(t *testing.T) {
	v := pml.NewCT_GroupShapeNonVisual()
	if v == nil {
		t.Errorf("pml.NewCT_GroupShapeNonVisual must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_GroupShapeNonVisual should validate: %s", err)
	}
}

func TestCT_GroupShapeNonVisualMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_GroupShapeNonVisual()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_GroupShapeNonVisual()
	xml.Unmarshal(buf, v2)
}
