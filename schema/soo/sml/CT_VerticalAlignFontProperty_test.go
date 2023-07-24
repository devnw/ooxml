package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_VerticalAlignFontPropertyConstructor(t *testing.T) {
	v := sml.NewCT_VerticalAlignFontProperty()
	if v == nil {
		t.Errorf("sml.NewCT_VerticalAlignFontProperty must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_VerticalAlignFontProperty should validate: %s", err)
	}
}

func TestCT_VerticalAlignFontPropertyMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_VerticalAlignFontProperty()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_VerticalAlignFontProperty()
	xml.Unmarshal(buf, v2)
}
