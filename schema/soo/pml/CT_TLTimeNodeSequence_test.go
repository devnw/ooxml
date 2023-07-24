package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TLTimeNodeSequenceConstructor(t *testing.T) {
	v := pml.NewCT_TLTimeNodeSequence()
	if v == nil {
		t.Errorf("pml.NewCT_TLTimeNodeSequence must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLTimeNodeSequence should validate: %s", err)
	}
}

func TestCT_TLTimeNodeSequenceMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLTimeNodeSequence()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLTimeNodeSequence()
	xml.Unmarshal(buf, v2)
}
