package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcCT_DiagramConstructor(t *testing.T) {
	v := vml.NewOfcCT_Diagram()
	if v == nil {
		t.Errorf("vml.NewOfcCT_Diagram must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcCT_Diagram should validate: %s", err)
	}
}

func TestOfcCT_DiagramMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcCT_Diagram()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcCT_Diagram()
	xml.Unmarshal(buf, v2)
}
