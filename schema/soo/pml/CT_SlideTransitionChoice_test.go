package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_SlideTransitionChoiceConstructor(t *testing.T) {
	v := pml.NewCT_SlideTransitionChoice()
	if v == nil {
		t.Errorf("pml.NewCT_SlideTransitionChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_SlideTransitionChoice should validate: %s", err)
	}
}

func TestCT_SlideTransitionChoiceMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_SlideTransitionChoice()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_SlideTransitionChoice()
	xml.Unmarshal(buf, v2)
}
