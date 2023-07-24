package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestAG_ImageAttributesConstructor(t *testing.T) {
	v := vml.NewAG_ImageAttributes()
	if v == nil {
		t.Errorf("vml.NewAG_ImageAttributes must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.AG_ImageAttributes should validate: %s", err)
	}
}

func TestAG_ImageAttributesMarshalUnmarshal(t *testing.T) {
	v := vml.NewAG_ImageAttributes()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewAG_ImageAttributes()
	xml.Unmarshal(buf, v2)
}
