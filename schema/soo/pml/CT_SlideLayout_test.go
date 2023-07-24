package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_SlideLayoutConstructor(t *testing.T) {
	v := pml.NewCT_SlideLayout()
	if v == nil {
		t.Errorf("pml.NewCT_SlideLayout must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_SlideLayout should validate: %s", err)
	}
}

func TestCT_SlideLayoutMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_SlideLayout()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_SlideLayout()
	xml.Unmarshal(buf, v2)
}
