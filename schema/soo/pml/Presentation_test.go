package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestPresentationConstructor(t *testing.T) {
	v := pml.NewPresentation()
	if v == nil {
		t.Errorf("pml.NewPresentation must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.Presentation should validate: %s", err)
	}
}

func TestPresentationMarshalUnmarshal(t *testing.T) {
	v := pml.NewPresentation()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewPresentation()
	xml.Unmarshal(buf, v2)
}
