package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestTextboxConstructor(t *testing.T) {
	v := vml.NewTextbox()
	if v == nil {
		t.Errorf("vml.NewTextbox must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.Textbox should validate: %s", err)
	}
}

func TestTextboxMarshalUnmarshal(t *testing.T) {
	v := vml.NewTextbox()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewTextbox()
	xml.Unmarshal(buf, v2)
}
