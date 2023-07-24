package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TLTimeTargetElementConstructor(t *testing.T) {
	v := pml.NewCT_TLTimeTargetElement()
	if v == nil {
		t.Errorf("pml.NewCT_TLTimeTargetElement must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLTimeTargetElement should validate: %s", err)
	}
}

func TestCT_TLTimeTargetElementMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLTimeTargetElement()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLTimeTargetElement()
	xml.Unmarshal(buf, v2)
}
