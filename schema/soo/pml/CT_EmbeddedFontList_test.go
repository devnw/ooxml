package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_EmbeddedFontListConstructor(t *testing.T) {
	v := pml.NewCT_EmbeddedFontList()
	if v == nil {
		t.Errorf("pml.NewCT_EmbeddedFontList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_EmbeddedFontList should validate: %s", err)
	}
}

func TestCT_EmbeddedFontListMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_EmbeddedFontList()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_EmbeddedFontList()
	xml.Unmarshal(buf, v2)
}
