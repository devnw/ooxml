package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcCT_CalloutConstructor(t *testing.T) {
	v := vml.NewOfcCT_Callout()
	if v == nil {
		t.Errorf("vml.NewOfcCT_Callout must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcCT_Callout should validate: %s", err)
	}
}

func TestOfcCT_CalloutMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcCT_Callout()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcCT_Callout()
	xml.Unmarshal(buf, v2)
}
