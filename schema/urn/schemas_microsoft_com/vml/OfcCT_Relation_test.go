package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcCT_RelationConstructor(t *testing.T) {
	v := vml.NewOfcCT_Relation()
	if v == nil {
		t.Errorf("vml.NewOfcCT_Relation must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcCT_Relation should validate: %s", err)
	}
}

func TestOfcCT_RelationMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcCT_Relation()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcCT_Relation()
	xml.Unmarshal(buf, v2)
}
