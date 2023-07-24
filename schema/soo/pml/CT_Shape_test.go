package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_ShapeConstructor(t *testing.T) {
	v := pml.NewCT_Shape()
	if v == nil {
		t.Errorf("pml.NewCT_Shape must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_Shape should validate: %s", err)
	}
}

func TestCT_ShapeMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_Shape()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_Shape()
	xml.Unmarshal(buf, v2)
}
