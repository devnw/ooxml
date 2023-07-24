package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TLTextTargetElementConstructor(t *testing.T) {
	v := pml.NewCT_TLTextTargetElement()
	if v == nil {
		t.Errorf("pml.NewCT_TLTextTargetElement must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLTextTargetElement should validate: %s", err)
	}
}

func TestCT_TLTextTargetElementMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLTextTargetElement()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLTextTargetElement()
	xml.Unmarshal(buf, v2)
}
