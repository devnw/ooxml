package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_IndexedColorsConstructor(t *testing.T) {
	v := sml.NewCT_IndexedColors()
	if v == nil {
		t.Errorf("sml.NewCT_IndexedColors must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_IndexedColors should validate: %s", err)
	}
}

func TestCT_IndexedColorsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_IndexedColors()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_IndexedColors()
	xml.Unmarshal(buf, v2)
}
