package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestCT_TextboxConstructor(t *testing.T) {
	v := vml.NewCT_Textbox()
	if v == nil {
		t.Errorf("vml.NewCT_Textbox must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.CT_Textbox should validate: %s", err)
	}
}

func TestCT_TextboxMarshalUnmarshal(t *testing.T) {
	v := vml.NewCT_Textbox()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewCT_Textbox()
	xml.Unmarshal(buf, v2)
}
