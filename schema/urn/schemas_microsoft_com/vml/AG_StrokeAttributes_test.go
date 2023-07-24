package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestAG_StrokeAttributesConstructor(t *testing.T) {
	v := vml.NewAG_StrokeAttributes()
	if v == nil {
		t.Errorf("vml.NewAG_StrokeAttributes must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.AG_StrokeAttributes should validate: %s", err)
	}
}

func TestAG_StrokeAttributesMarshalUnmarshal(t *testing.T) {
	v := vml.NewAG_StrokeAttributes()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewAG_StrokeAttributes()
	xml.Unmarshal(buf, v2)
}
