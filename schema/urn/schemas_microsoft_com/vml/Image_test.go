package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestImageConstructor(t *testing.T) {
	v := vml.NewImage()
	if v == nil {
		t.Errorf("vml.NewImage must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.Image should validate: %s", err)
	}
}

func TestImageMarshalUnmarshal(t *testing.T) {
	v := vml.NewImage()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewImage()
	xml.Unmarshal(buf, v2)
}
