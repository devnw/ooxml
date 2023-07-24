package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcCT_RegroupTableConstructor(t *testing.T) {
	v := vml.NewOfcCT_RegroupTable()
	if v == nil {
		t.Errorf("vml.NewOfcCT_RegroupTable must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcCT_RegroupTable should validate: %s", err)
	}
}

func TestOfcCT_RegroupTableMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcCT_RegroupTable()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcCT_RegroupTable()
	xml.Unmarshal(buf, v2)
}
