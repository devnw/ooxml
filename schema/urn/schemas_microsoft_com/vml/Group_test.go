package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestGroupConstructor(t *testing.T) {
	v := vml.NewGroup()
	if v == nil {
		t.Errorf("vml.NewGroup must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.Group should validate: %s", err)
	}
}

func TestGroupMarshalUnmarshal(t *testing.T) {
	v := vml.NewGroup()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewGroup()
	xml.Unmarshal(buf, v2)
}
