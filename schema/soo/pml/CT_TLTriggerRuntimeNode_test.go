package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TLTriggerRuntimeNodeConstructor(t *testing.T) {
	v := pml.NewCT_TLTriggerRuntimeNode()
	if v == nil {
		t.Errorf("pml.NewCT_TLTriggerRuntimeNode must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLTriggerRuntimeNode should validate: %s", err)
	}
}

func TestCT_TLTriggerRuntimeNodeMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLTriggerRuntimeNode()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLTriggerRuntimeNode()
	xml.Unmarshal(buf, v2)
}
