package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestAG_ShapeAttributesConstructor(t *testing.T) {
	v := vml.NewAG_ShapeAttributes()
	if v == nil {
		t.Errorf("vml.NewAG_ShapeAttributes must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.AG_ShapeAttributes should validate: %s", err)
	}
}

func TestAG_ShapeAttributesMarshalUnmarshal(t *testing.T) {
	v := vml.NewAG_ShapeAttributes()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewAG_ShapeAttributes()
	xml.Unmarshal(buf, v2)
}
