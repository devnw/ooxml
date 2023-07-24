package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestFillConstructor(t *testing.T) {
	v := vml.NewFill()
	if v == nil {
		t.Errorf("vml.NewFill must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.Fill should validate: %s", err)
	}
}

func TestFillMarshalUnmarshal(t *testing.T) {
	v := vml.NewFill()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewFill()
	xml.Unmarshal(buf, v2)
}
