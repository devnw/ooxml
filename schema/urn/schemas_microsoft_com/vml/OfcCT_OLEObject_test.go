package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcCT_OLEObjectConstructor(t *testing.T) {
	v := vml.NewOfcCT_OLEObject()
	if v == nil {
		t.Errorf("vml.NewOfcCT_OLEObject must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcCT_OLEObject should validate: %s", err)
	}
}

func TestOfcCT_OLEObjectMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcCT_OLEObject()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcCT_OLEObject()
	xml.Unmarshal(buf, v2)
}
