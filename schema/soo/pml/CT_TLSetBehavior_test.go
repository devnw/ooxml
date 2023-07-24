package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TLSetBehaviorConstructor(t *testing.T) {
	v := pml.NewCT_TLSetBehavior()
	if v == nil {
		t.Errorf("pml.NewCT_TLSetBehavior must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLSetBehavior should validate: %s", err)
	}
}

func TestCT_TLSetBehaviorMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLSetBehavior()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLSetBehavior()
	xml.Unmarshal(buf, v2)
}
