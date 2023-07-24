package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcInkConstructor(t *testing.T) {
	v := vml.NewOfcInk()
	if v == nil {
		t.Errorf("vml.NewOfcInk must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcInk should validate: %s", err)
	}
}

func TestOfcInkMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcInk()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcInk()
	xml.Unmarshal(buf, v2)
}
