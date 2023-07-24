package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestCT_LineConstructor(t *testing.T) {
	v := vml.NewCT_Line()
	if v == nil {
		t.Errorf("vml.NewCT_Line must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.CT_Line should validate: %s", err)
	}
}

func TestCT_LineMarshalUnmarshal(t *testing.T) {
	v := vml.NewCT_Line()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewCT_Line()
	xml.Unmarshal(buf, v2)
}
