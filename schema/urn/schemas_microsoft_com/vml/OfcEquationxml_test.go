package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcEquationxmlConstructor(t *testing.T) {
	v := vml.NewOfcEquationxml()
	if v == nil {
		t.Errorf("vml.NewOfcEquationxml must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcEquationxml should validate: %s", err)
	}
}

func TestOfcEquationxmlMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcEquationxml()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcEquationxml()
	xml.Unmarshal(buf, v2)
}
