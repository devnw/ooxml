package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestEG_ChildSlideConstructor(t *testing.T) {
	v := pml.NewEG_ChildSlide()
	if v == nil {
		t.Errorf("pml.NewEG_ChildSlide must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.EG_ChildSlide should validate: %s", err)
	}
}

func TestEG_ChildSlideMarshalUnmarshal(t *testing.T) {
	v := pml.NewEG_ChildSlide()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewEG_ChildSlide()
	xml.Unmarshal(buf, v2)
}
