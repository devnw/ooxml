package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestCT_ShadowConstructor(t *testing.T) {
	v := vml.NewCT_Shadow()
	if v == nil {
		t.Errorf("vml.NewCT_Shadow must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.CT_Shadow should validate: %s", err)
	}
}

func TestCT_ShadowMarshalUnmarshal(t *testing.T) {
	v := vml.NewCT_Shadow()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewCT_Shadow()
	xml.Unmarshal(buf, v2)
}
