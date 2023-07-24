package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestTextpathConstructor(t *testing.T) {
	v := vml.NewTextpath()
	if v == nil {
		t.Errorf("vml.NewTextpath must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.Textpath should validate: %s", err)
	}
}

func TestTextpathMarshalUnmarshal(t *testing.T) {
	v := vml.NewTextpath()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewTextpath()
	xml.Unmarshal(buf, v2)
}
