package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestCT_ShapeConstructor(t *testing.T) {
	v := vml.NewCT_Shape()
	if v == nil {
		t.Errorf("vml.NewCT_Shape must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.CT_Shape should validate: %s", err)
	}
}

func TestCT_ShapeMarshalUnmarshal(t *testing.T) {
	v := vml.NewCT_Shape()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewCT_Shape()
	xml.Unmarshal(buf, v2)
}
