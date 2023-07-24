package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestShadowConstructor(t *testing.T) {
	v := vml.NewShadow()
	if v == nil {
		t.Errorf("vml.NewShadow must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.Shadow should validate: %s", err)
	}
}

func TestShadowMarshalUnmarshal(t *testing.T) {
	v := vml.NewShadow()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewShadow()
	xml.Unmarshal(buf, v2)
}
