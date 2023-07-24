package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_CornerDirectionTransitionConstructor(t *testing.T) {
	v := pml.NewCT_CornerDirectionTransition()
	if v == nil {
		t.Errorf("pml.NewCT_CornerDirectionTransition must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_CornerDirectionTransition should validate: %s", err)
	}
}

func TestCT_CornerDirectionTransitionMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_CornerDirectionTransition()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_CornerDirectionTransition()
	xml.Unmarshal(buf, v2)
}
