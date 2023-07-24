package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_EightDirectionTransitionConstructor(t *testing.T) {
	v := pml.NewCT_EightDirectionTransition()
	if v == nil {
		t.Errorf("pml.NewCT_EightDirectionTransition must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_EightDirectionTransition should validate: %s", err)
	}
}

func TestCT_EightDirectionTransitionMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_EightDirectionTransition()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_EightDirectionTransition()
	xml.Unmarshal(buf, v2)
}
