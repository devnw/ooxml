package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcCT_RConstructor(t *testing.T) {
	v := vml.NewOfcCT_R()
	if v == nil {
		t.Errorf("vml.NewOfcCT_R must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcCT_R should validate: %s", err)
	}
}

func TestOfcCT_RMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcCT_R()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcCT_R()
	xml.Unmarshal(buf, v2)
}
