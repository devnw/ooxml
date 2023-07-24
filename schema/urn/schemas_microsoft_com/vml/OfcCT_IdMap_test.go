package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcCT_IdMapConstructor(t *testing.T) {
	v := vml.NewOfcCT_IdMap()
	if v == nil {
		t.Errorf("vml.NewOfcCT_IdMap must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcCT_IdMap should validate: %s", err)
	}
}

func TestOfcCT_IdMapMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcCT_IdMap()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcCT_IdMap()
	xml.Unmarshal(buf, v2)
}
