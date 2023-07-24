package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcCT_ClipPathConstructor(t *testing.T) {
	v := vml.NewOfcCT_ClipPath()
	if v == nil {
		t.Errorf("vml.NewOfcCT_ClipPath must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcCT_ClipPath should validate: %s", err)
	}
}

func TestOfcCT_ClipPathMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcCT_ClipPath()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcCT_ClipPath()
	xml.Unmarshal(buf, v2)
}
