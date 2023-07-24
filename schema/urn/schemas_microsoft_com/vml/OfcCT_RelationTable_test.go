package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcCT_RelationTableConstructor(t *testing.T) {
	v := vml.NewOfcCT_RelationTable()
	if v == nil {
		t.Errorf("vml.NewOfcCT_RelationTable must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcCT_RelationTable should validate: %s", err)
	}
}

func TestOfcCT_RelationTableMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcCT_RelationTable()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcCT_RelationTable()
	xml.Unmarshal(buf, v2)
}
