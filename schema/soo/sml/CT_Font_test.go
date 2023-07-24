package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_FontConstructor(t *testing.T) {
	v := sml.NewCT_Font()
	if v == nil {
		t.Errorf("sml.NewCT_Font must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Font should validate: %s", err)
	}
}

func TestCT_FontMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Font()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Font()
	xml.Unmarshal(buf, v2)
}
