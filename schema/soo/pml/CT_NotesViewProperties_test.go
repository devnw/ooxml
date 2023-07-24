package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_NotesViewPropertiesConstructor(t *testing.T) {
	v := pml.NewCT_NotesViewProperties()
	if v == nil {
		t.Errorf("pml.NewCT_NotesViewProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_NotesViewProperties should validate: %s", err)
	}
}

func TestCT_NotesViewPropertiesMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_NotesViewProperties()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_NotesViewProperties()
	xml.Unmarshal(buf, v2)
}
