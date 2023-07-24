package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestPresentationPrConstructor(t *testing.T) {
	v := pml.NewPresentationPr()
	if v == nil {
		t.Errorf("pml.NewPresentationPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.PresentationPr should validate: %s", err)
	}
}

func TestPresentationPrMarshalUnmarshal(t *testing.T) {
	v := pml.NewPresentationPr()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewPresentationPr()
	xml.Unmarshal(buf, v2)
}
