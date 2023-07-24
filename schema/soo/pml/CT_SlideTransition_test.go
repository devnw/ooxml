package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_SlideTransitionConstructor(t *testing.T) {
	v := pml.NewCT_SlideTransition()
	if v == nil {
		t.Errorf("pml.NewCT_SlideTransition must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_SlideTransition should validate: %s", err)
	}
}

func TestCT_SlideTransitionMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_SlideTransition()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_SlideTransition()
	xml.Unmarshal(buf, v2)
}
