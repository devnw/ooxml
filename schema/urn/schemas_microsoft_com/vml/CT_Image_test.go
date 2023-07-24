package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestCT_ImageConstructor(t *testing.T) {
	v := vml.NewCT_Image()
	if v == nil {
		t.Errorf("vml.NewCT_Image must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.CT_Image should validate: %s", err)
	}
}

func TestCT_ImageMarshalUnmarshal(t *testing.T) {
	v := vml.NewCT_Image()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewCT_Image()
	xml.Unmarshal(buf, v2)
}
