package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TLBehaviorAttributeNameListConstructor(t *testing.T) {
	v := pml.NewCT_TLBehaviorAttributeNameList()
	if v == nil {
		t.Errorf("pml.NewCT_TLBehaviorAttributeNameList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLBehaviorAttributeNameList should validate: %s", err)
	}
}

func TestCT_TLBehaviorAttributeNameListMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLBehaviorAttributeNameList()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLBehaviorAttributeNameList()
	xml.Unmarshal(buf, v2)
}
