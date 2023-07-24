package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcCT_ColorMenuConstructor(t *testing.T) {
	v := vml.NewOfcCT_ColorMenu()
	if v == nil {
		t.Errorf("vml.NewOfcCT_ColorMenu must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcCT_ColorMenu should validate: %s", err)
	}
}

func TestOfcCT_ColorMenuMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcCT_ColorMenu()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcCT_ColorMenu()
	xml.Unmarshal(buf, v2)
}
