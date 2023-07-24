package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcCT_EquationXmlConstructor(t *testing.T) {
	v := vml.NewOfcCT_EquationXml()
	if v == nil {
		t.Errorf("vml.NewOfcCT_EquationXml must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcCT_EquationXml should validate: %s", err)
	}
}

func TestOfcCT_EquationXmlMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcCT_EquationXml()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcCT_EquationXml()
	xml.Unmarshal(buf, v2)
}
