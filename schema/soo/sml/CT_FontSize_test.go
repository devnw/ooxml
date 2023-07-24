package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_FontSizeConstructor(t *testing.T) {
	v := sml.NewCT_FontSize()
	if v == nil {
		t.Errorf("sml.NewCT_FontSize must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_FontSize should validate: %s", err)
	}
}

func TestCT_FontSizeMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_FontSize()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_FontSize()
	xml.Unmarshal(buf, v2)
}
