package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_EmbeddedFontListEntryConstructor(t *testing.T) {
	v := pml.NewCT_EmbeddedFontListEntry()
	if v == nil {
		t.Errorf("pml.NewCT_EmbeddedFontListEntry must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_EmbeddedFontListEntry should validate: %s", err)
	}
}

func TestCT_EmbeddedFontListEntryMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_EmbeddedFontListEntry()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_EmbeddedFontListEntry()
	xml.Unmarshal(buf, v2)
}
