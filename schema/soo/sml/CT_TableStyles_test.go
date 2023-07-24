package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_TableStylesConstructor(t *testing.T) {
	v := sml.NewCT_TableStyles()
	if v == nil {
		t.Errorf("sml.NewCT_TableStyles must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_TableStyles should validate: %s", err)
	}
}

func TestCT_TableStylesMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_TableStyles()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_TableStyles()
	xml.Unmarshal(buf, v2)
}
