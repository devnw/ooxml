package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcCT_SignatureLineConstructor(t *testing.T) {
	v := vml.NewOfcCT_SignatureLine()
	if v == nil {
		t.Errorf("vml.NewOfcCT_SignatureLine must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcCT_SignatureLine should validate: %s", err)
	}
}

func TestOfcCT_SignatureLineMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcCT_SignatureLine()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcCT_SignatureLine()
	xml.Unmarshal(buf, v2)
}
