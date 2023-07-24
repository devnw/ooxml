package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcLeftConstructor(t *testing.T) {
	v := vml.NewOfcLeft()
	if v == nil {
		t.Errorf("vml.NewOfcLeft must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcLeft should validate: %s", err)
	}
}

func TestOfcLeftMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcLeft()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcLeft()
	xml.Unmarshal(buf, v2)
}
