package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_NotesMasterIdListEntryConstructor(t *testing.T) {
	v := pml.NewCT_NotesMasterIdListEntry()
	if v == nil {
		t.Errorf("pml.NewCT_NotesMasterIdListEntry must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_NotesMasterIdListEntry should validate: %s", err)
	}
}

func TestCT_NotesMasterIdListEntryMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_NotesMasterIdListEntry()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_NotesMasterIdListEntry()
	xml.Unmarshal(buf, v2)
}
