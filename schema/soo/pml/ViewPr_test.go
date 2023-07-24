package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestViewPrConstructor(t *testing.T) {
	v := pml.NewViewPr()
	if v == nil {
		t.Errorf("pml.NewViewPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.ViewPr should validate: %s", err)
	}
}

func TestViewPrMarshalUnmarshal(t *testing.T) {
	v := pml.NewViewPr()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewViewPr()
	xml.Unmarshal(buf, v2)
}
