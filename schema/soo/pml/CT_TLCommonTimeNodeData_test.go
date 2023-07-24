package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TLCommonTimeNodeDataConstructor(t *testing.T) {
	v := pml.NewCT_TLCommonTimeNodeData()
	if v == nil {
		t.Errorf("pml.NewCT_TLCommonTimeNodeData must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLCommonTimeNodeData should validate: %s", err)
	}
}

func TestCT_TLCommonTimeNodeDataMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLCommonTimeNodeData()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLCommonTimeNodeData()
	xml.Unmarshal(buf, v2)
}
