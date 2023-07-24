package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestLineConstructor(t *testing.T) {
	v := vml.NewLine()
	if v == nil {
		t.Errorf("vml.NewLine must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.Line should validate: %s", err)
	}
}

func TestLineMarshalUnmarshal(t *testing.T) {
	v := vml.NewLine()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewLine()
	xml.Unmarshal(buf, v2)
}
