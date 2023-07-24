package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestNotesConstructor(t *testing.T) {
	v := pml.NewNotes()
	if v == nil {
		t.Errorf("pml.NewNotes must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.Notes should validate: %s", err)
	}
}

func TestNotesMarshalUnmarshal(t *testing.T) {
	v := pml.NewNotes()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewNotes()
	xml.Unmarshal(buf, v2)
}
