package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_NotesSlideConstructor(t *testing.T) {
	v := pml.NewCT_NotesSlide()
	if v == nil {
		t.Errorf("pml.NewCT_NotesSlide must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_NotesSlide should validate: %s", err)
	}
}

func TestCT_NotesSlideMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_NotesSlide()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_NotesSlide()
	xml.Unmarshal(buf, v2)
}
