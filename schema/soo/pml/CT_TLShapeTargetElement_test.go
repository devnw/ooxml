package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TLShapeTargetElementConstructor(t *testing.T) {
	v := pml.NewCT_TLShapeTargetElement()
	if v == nil {
		t.Errorf("pml.NewCT_TLShapeTargetElement must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLShapeTargetElement should validate: %s", err)
	}
}

func TestCT_TLShapeTargetElementMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLShapeTargetElement()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLShapeTargetElement()
	xml.Unmarshal(buf, v2)
}
