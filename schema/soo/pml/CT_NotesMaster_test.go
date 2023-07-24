package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_NotesMasterConstructor(t *testing.T) {
	v := pml.NewCT_NotesMaster()
	if v == nil {
		t.Errorf("pml.NewCT_NotesMaster must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_NotesMaster should validate: %s", err)
	}
}

func TestCT_NotesMasterMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_NotesMaster()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_NotesMaster()
	xml.Unmarshal(buf, v2)
}
