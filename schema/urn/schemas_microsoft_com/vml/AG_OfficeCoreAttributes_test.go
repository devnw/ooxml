package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestAG_OfficeCoreAttributesConstructor(t *testing.T) {
	v := vml.NewAG_OfficeCoreAttributes()
	if v == nil {
		t.Errorf("vml.NewAG_OfficeCoreAttributes must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.AG_OfficeCoreAttributes should validate: %s", err)
	}
}

func TestAG_OfficeCoreAttributesMarshalUnmarshal(t *testing.T) {
	v := vml.NewAG_OfficeCoreAttributes()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewAG_OfficeCoreAttributes()
	xml.Unmarshal(buf, v2)
}
