package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_PresentationConstructor(t *testing.T) {
	v := pml.NewCT_Presentation()
	if v == nil {
		t.Errorf("pml.NewCT_Presentation must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_Presentation should validate: %s", err)
	}
}

func TestCT_PresentationMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_Presentation()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_Presentation()
	xml.Unmarshal(buf, v2)
}
