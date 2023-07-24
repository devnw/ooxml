package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_FontFamilyConstructor(t *testing.T) {
	v := sml.NewCT_FontFamily()
	if v == nil {
		t.Errorf("sml.NewCT_FontFamily must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_FontFamily should validate: %s", err)
	}
}

func TestCT_FontFamilyMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_FontFamily()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_FontFamily()
	xml.Unmarshal(buf, v2)
}
