package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_MRUColorsConstructor(t *testing.T) {
	v := sml.NewCT_MRUColors()
	if v == nil {
		t.Errorf("sml.NewCT_MRUColors must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_MRUColors should validate: %s", err)
	}
}

func TestCT_MRUColorsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_MRUColors()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_MRUColors()
	xml.Unmarshal(buf, v2)
}
