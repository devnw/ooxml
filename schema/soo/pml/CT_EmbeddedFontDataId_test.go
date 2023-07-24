package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_EmbeddedFontDataIdConstructor(t *testing.T) {
	v := pml.NewCT_EmbeddedFontDataId()
	if v == nil {
		t.Errorf("pml.NewCT_EmbeddedFontDataId must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_EmbeddedFontDataId should validate: %s", err)
	}
}

func TestCT_EmbeddedFontDataIdMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_EmbeddedFontDataId()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_EmbeddedFontDataId()
	xml.Unmarshal(buf, v2)
}
