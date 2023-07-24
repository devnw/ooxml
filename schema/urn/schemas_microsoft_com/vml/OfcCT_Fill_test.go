package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcCT_FillConstructor(t *testing.T) {
	v := vml.NewOfcCT_Fill()
	if v == nil {
		t.Errorf("vml.NewOfcCT_Fill must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcCT_Fill should validate: %s", err)
	}
}

func TestOfcCT_FillMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcCT_Fill()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcCT_Fill()
	xml.Unmarshal(buf, v2)
}
