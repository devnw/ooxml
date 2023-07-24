package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_GroupShapeConstructor(t *testing.T) {
	v := pml.NewCT_GroupShape()
	if v == nil {
		t.Errorf("pml.NewCT_GroupShape must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_GroupShape should validate: %s", err)
	}
}

func TestCT_GroupShapeMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_GroupShape()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_GroupShape()
	xml.Unmarshal(buf, v2)
}
