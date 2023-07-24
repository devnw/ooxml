package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_SlideConstructor(t *testing.T) {
	v := pml.NewCT_Slide()
	if v == nil {
		t.Errorf("pml.NewCT_Slide must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_Slide should validate: %s", err)
	}
}

func TestCT_SlideMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_Slide()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_Slide()
	xml.Unmarshal(buf, v2)
}
