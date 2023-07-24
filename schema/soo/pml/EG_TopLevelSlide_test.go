package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestEG_TopLevelSlideConstructor(t *testing.T) {
	v := pml.NewEG_TopLevelSlide()
	if v == nil {
		t.Errorf("pml.NewEG_TopLevelSlide must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.EG_TopLevelSlide should validate: %s", err)
	}
}

func TestEG_TopLevelSlideMarshalUnmarshal(t *testing.T) {
	v := pml.NewEG_TopLevelSlide()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewEG_TopLevelSlide()
	xml.Unmarshal(buf, v2)
}
