package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcCT_ShapeLayoutConstructor(t *testing.T) {
	v := vml.NewOfcCT_ShapeLayout()
	if v == nil {
		t.Errorf("vml.NewOfcCT_ShapeLayout must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcCT_ShapeLayout should validate: %s", err)
	}
}

func TestOfcCT_ShapeLayoutMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcCT_ShapeLayout()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcCT_ShapeLayout()
	xml.Unmarshal(buf, v2)
}
