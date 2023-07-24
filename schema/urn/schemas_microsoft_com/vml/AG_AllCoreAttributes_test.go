package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestAG_AllCoreAttributesConstructor(t *testing.T) {
	v := vml.NewAG_AllCoreAttributes()
	if v == nil {
		t.Errorf("vml.NewAG_AllCoreAttributes must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.AG_AllCoreAttributes should validate: %s", err)
	}
}

func TestAG_AllCoreAttributesMarshalUnmarshal(t *testing.T) {
	v := vml.NewAG_AllCoreAttributes()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewAG_AllCoreAttributes()
	xml.Unmarshal(buf, v2)
}
