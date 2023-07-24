package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TLCommonBehaviorDataConstructor(t *testing.T) {
	v := pml.NewCT_TLCommonBehaviorData()
	if v == nil {
		t.Errorf("pml.NewCT_TLCommonBehaviorData must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLCommonBehaviorData should validate: %s", err)
	}
}

func TestCT_TLCommonBehaviorDataMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLCommonBehaviorData()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLCommonBehaviorData()
	xml.Unmarshal(buf, v2)
}
