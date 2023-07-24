package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestImagedataConstructor(t *testing.T) {
	v := vml.NewImagedata()
	if v == nil {
		t.Errorf("vml.NewImagedata must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.Imagedata should validate: %s", err)
	}
}

func TestImagedataMarshalUnmarshal(t *testing.T) {
	v := vml.NewImagedata()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewImagedata()
	xml.Unmarshal(buf, v2)
}
