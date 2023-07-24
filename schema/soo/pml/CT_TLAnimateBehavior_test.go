package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TLAnimateBehaviorConstructor(t *testing.T) {
	v := pml.NewCT_TLAnimateBehavior()
	if v == nil {
		t.Errorf("pml.NewCT_TLAnimateBehavior must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLAnimateBehavior should validate: %s", err)
	}
}

func TestCT_TLAnimateBehaviorMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLAnimateBehavior()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLAnimateBehavior()
	xml.Unmarshal(buf, v2)
}
