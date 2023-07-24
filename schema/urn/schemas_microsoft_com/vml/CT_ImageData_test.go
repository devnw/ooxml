package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestCT_ImageDataConstructor(t *testing.T) {
	v := vml.NewCT_ImageData()
	if v == nil {
		t.Errorf("vml.NewCT_ImageData must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.CT_ImageData should validate: %s", err)
	}
}

func TestCT_ImageDataMarshalUnmarshal(t *testing.T) {
	v := vml.NewCT_ImageData()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewCT_ImageData()
	xml.Unmarshal(buf, v2)
}
