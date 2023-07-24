package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcShapedefaultsConstructor(t *testing.T) {
	v := vml.NewOfcShapedefaults()
	if v == nil {
		t.Errorf("vml.NewOfcShapedefaults must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcShapedefaults should validate: %s", err)
	}
}

func TestOfcShapedefaultsMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcShapedefaults()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcShapedefaults()
	xml.Unmarshal(buf, v2)
}
