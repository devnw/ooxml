package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_SlideSizeConstructor(t *testing.T) {
	v := pml.NewCT_SlideSize()
	if v == nil {
		t.Errorf("pml.NewCT_SlideSize must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_SlideSize should validate: %s", err)
	}
}

func TestCT_SlideSizeMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_SlideSize()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_SlideSize()
	xml.Unmarshal(buf, v2)
}
