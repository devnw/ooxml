package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestAG_ChildSlideConstructor(t *testing.T) {
	v := pml.NewAG_ChildSlide()
	if v == nil {
		t.Errorf("pml.NewAG_ChildSlide must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.AG_ChildSlide should validate: %s", err)
	}
}

func TestAG_ChildSlideMarshalUnmarshal(t *testing.T) {
	v := pml.NewAG_ChildSlide()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewAG_ChildSlide()
	xml.Unmarshal(buf, v2)
}
