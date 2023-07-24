package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_InOutTransitionConstructor(t *testing.T) {
	v := pml.NewCT_InOutTransition()
	if v == nil {
		t.Errorf("pml.NewCT_InOutTransition must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_InOutTransition should validate: %s", err)
	}
}

func TestCT_InOutTransitionMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_InOutTransition()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_InOutTransition()
	xml.Unmarshal(buf, v2)
}
