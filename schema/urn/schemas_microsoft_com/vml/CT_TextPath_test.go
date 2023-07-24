package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestCT_TextPathConstructor(t *testing.T) {
	v := vml.NewCT_TextPath()
	if v == nil {
		t.Errorf("vml.NewCT_TextPath must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.CT_TextPath should validate: %s", err)
	}
}

func TestCT_TextPathMarshalUnmarshal(t *testing.T) {
	v := vml.NewCT_TextPath()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewCT_TextPath()
	xml.Unmarshal(buf, v2)
}
