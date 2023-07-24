package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcCalloutConstructor(t *testing.T) {
	v := vml.NewOfcCallout()
	if v == nil {
		t.Errorf("vml.NewOfcCallout must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcCallout should validate: %s", err)
	}
}

func TestOfcCalloutMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcCallout()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcCallout()
	xml.Unmarshal(buf, v2)
}
