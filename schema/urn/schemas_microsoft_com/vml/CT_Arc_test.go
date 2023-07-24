package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestCT_ArcConstructor(t *testing.T) {
	v := vml.NewCT_Arc()
	if v == nil {
		t.Errorf("vml.NewCT_Arc must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.CT_Arc should validate: %s", err)
	}
}

func TestCT_ArcMarshalUnmarshal(t *testing.T) {
	v := vml.NewCT_Arc()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewCT_Arc()
	xml.Unmarshal(buf, v2)
}
