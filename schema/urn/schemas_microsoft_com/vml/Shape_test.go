package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestShapeConstructor(t *testing.T) {
	v := vml.NewShape()
	if v == nil {
		t.Errorf("vml.NewShape must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.Shape should validate: %s", err)
	}
}

func TestShapeMarshalUnmarshal(t *testing.T) {
	v := vml.NewShape()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewShape()
	xml.Unmarshal(buf, v2)
}
