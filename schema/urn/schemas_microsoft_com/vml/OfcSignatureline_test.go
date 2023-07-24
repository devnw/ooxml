package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcSignaturelineConstructor(t *testing.T) {
	v := vml.NewOfcSignatureline()
	if v == nil {
		t.Errorf("vml.NewOfcSignatureline must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcSignatureline should validate: %s", err)
	}
}

func TestOfcSignaturelineMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcSignatureline()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcSignatureline()
	xml.Unmarshal(buf, v2)
}
