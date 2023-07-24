package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestCT_ShapetypeConstructor(t *testing.T) {
	v := vml.NewCT_Shapetype()
	if v == nil {
		t.Errorf("vml.NewCT_Shapetype must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.CT_Shapetype should validate: %s", err)
	}
}

func TestCT_ShapetypeMarshalUnmarshal(t *testing.T) {
	v := vml.NewCT_Shapetype()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewCT_Shapetype()
	xml.Unmarshal(buf, v2)
}
