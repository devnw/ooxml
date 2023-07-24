package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ServerFormatConstructor(t *testing.T) {
	v := sml.NewCT_ServerFormat()
	if v == nil {
		t.Errorf("sml.NewCT_ServerFormat must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_ServerFormat should validate: %s", err)
	}
}

func TestCT_ServerFormatMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_ServerFormat()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_ServerFormat()
	xml.Unmarshal(buf, v2)
}
