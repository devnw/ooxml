package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcFillConstructor(t *testing.T) {
	v := vml.NewOfcFill()
	if v == nil {
		t.Errorf("vml.NewOfcFill must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcFill should validate: %s", err)
	}
}

func TestOfcFillMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcFill()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcFill()
	xml.Unmarshal(buf, v2)
}
