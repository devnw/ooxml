package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcCT_ShapeDefaultsConstructor(t *testing.T) {
	v := vml.NewOfcCT_ShapeDefaults()
	if v == nil {
		t.Errorf("vml.NewOfcCT_ShapeDefaults must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcCT_ShapeDefaults should validate: %s", err)
	}
}

func TestOfcCT_ShapeDefaultsMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcCT_ShapeDefaults()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcCT_ShapeDefaults()
	xml.Unmarshal(buf, v2)
}
