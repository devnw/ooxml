package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcCT_ColorMruConstructor(t *testing.T) {
	v := vml.NewOfcCT_ColorMru()
	if v == nil {
		t.Errorf("vml.NewOfcCT_ColorMru must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcCT_ColorMru should validate: %s", err)
	}
}

func TestOfcCT_ColorMruMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcCT_ColorMru()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcCT_ColorMru()
	xml.Unmarshal(buf, v2)
}
