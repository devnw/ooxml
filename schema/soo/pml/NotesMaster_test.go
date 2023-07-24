package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestNotesMasterConstructor(t *testing.T) {
	v := pml.NewNotesMaster()
	if v == nil {
		t.Errorf("pml.NewNotesMaster must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.NotesMaster should validate: %s", err)
	}
}

func TestNotesMasterMarshalUnmarshal(t *testing.T) {
	v := pml.NewNotesMaster()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewNotesMaster()
	xml.Unmarshal(buf, v2)
}
