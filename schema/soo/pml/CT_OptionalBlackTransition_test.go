package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_OptionalBlackTransitionConstructor(t *testing.T) {
	v := pml.NewCT_OptionalBlackTransition()
	if v == nil {
		t.Errorf("pml.NewCT_OptionalBlackTransition must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_OptionalBlackTransition should validate: %s", err)
	}
}

func TestCT_OptionalBlackTransitionMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_OptionalBlackTransition()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_OptionalBlackTransition()
	xml.Unmarshal(buf, v2)
}
