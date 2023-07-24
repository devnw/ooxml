package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TLSubShapeIdConstructor(t *testing.T) {
	v := pml.NewCT_TLSubShapeId()
	if v == nil {
		t.Errorf("pml.NewCT_TLSubShapeId must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLSubShapeId should validate: %s", err)
	}
}

func TestCT_TLSubShapeIdMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLSubShapeId()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLSubShapeId()
	xml.Unmarshal(buf, v2)
}
