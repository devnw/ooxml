package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcShapelayoutConstructor(t *testing.T) {
	v := vml.NewOfcShapelayout()
	if v == nil {
		t.Errorf("vml.NewOfcShapelayout must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcShapelayout should validate: %s", err)
	}
}

func TestOfcShapelayoutMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcShapelayout()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcShapelayout()
	xml.Unmarshal(buf, v2)
}
