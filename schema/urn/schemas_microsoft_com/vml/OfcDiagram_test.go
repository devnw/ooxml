package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcDiagramConstructor(t *testing.T) {
	v := vml.NewOfcDiagram()
	if v == nil {
		t.Errorf("vml.NewOfcDiagram must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcDiagram should validate: %s", err)
	}
}

func TestOfcDiagramMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcDiagram()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcDiagram()
	xml.Unmarshal(buf, v2)
}
