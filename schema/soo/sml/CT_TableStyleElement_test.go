package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_TableStyleElementConstructor(t *testing.T) {
	v := sml.NewCT_TableStyleElement()
	if v == nil {
		t.Errorf("sml.NewCT_TableStyleElement must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_TableStyleElement should validate: %s", err)
	}
}

func TestCT_TableStyleElementMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_TableStyleElement()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_TableStyleElement()
	xml.Unmarshal(buf, v2)
}
