package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TLTriggerTimeNodeIDConstructor(t *testing.T) {
	v := pml.NewCT_TLTriggerTimeNodeID()
	if v == nil {
		t.Errorf("pml.NewCT_TLTriggerTimeNodeID must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLTriggerTimeNodeID should validate: %s", err)
	}
}

func TestCT_TLTriggerTimeNodeIDMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLTriggerTimeNodeID()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLTriggerTimeNodeID()
	xml.Unmarshal(buf, v2)
}
