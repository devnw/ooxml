package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcComplexConstructor(t *testing.T) {
	v := vml.NewOfcComplex()
	if v == nil {
		t.Errorf("vml.NewOfcComplex must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcComplex should validate: %s", err)
	}
}

func TestOfcComplexMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcComplex()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcComplex()
	xml.Unmarshal(buf, v2)
}
