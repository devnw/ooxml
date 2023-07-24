package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestHandlesConstructor(t *testing.T) {
	v := vml.NewHandles()
	if v == nil {
		t.Errorf("vml.NewHandles must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.Handles should validate: %s", err)
	}
}

func TestHandlesMarshalUnmarshal(t *testing.T) {
	v := vml.NewHandles()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewHandles()
	xml.Unmarshal(buf, v2)
}
