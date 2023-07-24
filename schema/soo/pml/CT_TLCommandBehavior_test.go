package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TLCommandBehaviorConstructor(t *testing.T) {
	v := pml.NewCT_TLCommandBehavior()
	if v == nil {
		t.Errorf("pml.NewCT_TLCommandBehavior must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLCommandBehavior should validate: %s", err)
	}
}

func TestCT_TLCommandBehaviorMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLCommandBehavior()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLCommandBehavior()
	xml.Unmarshal(buf, v2)
}
