package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TLCommonMediaNodeDataConstructor(t *testing.T) {
	v := pml.NewCT_TLCommonMediaNodeData()
	if v == nil {
		t.Errorf("pml.NewCT_TLCommonMediaNodeData must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLCommonMediaNodeData should validate: %s", err)
	}
}

func TestCT_TLCommonMediaNodeDataMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLCommonMediaNodeData()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLCommonMediaNodeData()
	xml.Unmarshal(buf, v2)
}
