package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestAG_AllShapeAttributesConstructor(t *testing.T) {
	v := vml.NewAG_AllShapeAttributes()
	if v == nil {
		t.Errorf("vml.NewAG_AllShapeAttributes must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.AG_AllShapeAttributes should validate: %s", err)
	}
}

func TestAG_AllShapeAttributesMarshalUnmarshal(t *testing.T) {
	v := vml.NewAG_AllShapeAttributes()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewAG_AllShapeAttributes()
	xml.Unmarshal(buf, v2)
}
