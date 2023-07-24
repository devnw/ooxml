package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_XStringElementConstructor(t *testing.T) {
	v := sml.NewCT_XStringElement()
	if v == nil {
		t.Errorf("sml.NewCT_XStringElement must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_XStringElement should validate: %s", err)
	}
}

func TestCT_XStringElementMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_XStringElement()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_XStringElement()
	xml.Unmarshal(buf, v2)
}
