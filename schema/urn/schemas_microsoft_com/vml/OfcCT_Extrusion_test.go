package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcCT_ExtrusionConstructor(t *testing.T) {
	v := vml.NewOfcCT_Extrusion()
	if v == nil {
		t.Errorf("vml.NewOfcCT_Extrusion must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcCT_Extrusion should validate: %s", err)
	}
}

func TestOfcCT_ExtrusionMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcCT_Extrusion()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcCT_Extrusion()
	xml.Unmarshal(buf, v2)
}
