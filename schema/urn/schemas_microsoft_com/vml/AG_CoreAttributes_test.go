package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestAG_CoreAttributesConstructor(t *testing.T) {
	v := vml.NewAG_CoreAttributes()
	if v == nil {
		t.Errorf("vml.NewAG_CoreAttributes must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.AG_CoreAttributes should validate: %s", err)
	}
}

func TestAG_CoreAttributesMarshalUnmarshal(t *testing.T) {
	v := vml.NewAG_CoreAttributes()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewAG_CoreAttributes()
	xml.Unmarshal(buf, v2)
}
