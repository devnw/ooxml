package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcCT_ComplexConstructor(t *testing.T) {
	v := vml.NewOfcCT_Complex()
	if v == nil {
		t.Errorf("vml.NewOfcCT_Complex must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcCT_Complex should validate: %s", err)
	}
}

func TestOfcCT_ComplexMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcCT_Complex()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcCT_Complex()
	xml.Unmarshal(buf, v2)
}
