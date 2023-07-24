package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TLTimeConditionListConstructor(t *testing.T) {
	v := pml.NewCT_TLTimeConditionList()
	if v == nil {
		t.Errorf("pml.NewCT_TLTimeConditionList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLTimeConditionList should validate: %s", err)
	}
}

func TestCT_TLTimeConditionListMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLTimeConditionList()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLTimeConditionList()
	xml.Unmarshal(buf, v2)
}
