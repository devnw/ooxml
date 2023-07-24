package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_NotesTextViewPropertiesConstructor(t *testing.T) {
	v := pml.NewCT_NotesTextViewProperties()
	if v == nil {
		t.Errorf("pml.NewCT_NotesTextViewProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_NotesTextViewProperties should validate: %s", err)
	}
}

func TestCT_NotesTextViewPropertiesMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_NotesTextViewProperties()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_NotesTextViewProperties()
	xml.Unmarshal(buf, v2)
}
