package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcOLEObjectConstructor(t *testing.T) {
	v := vml.NewOfcOLEObject()
	if v == nil {
		t.Errorf("vml.NewOfcOLEObject must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcOLEObject should validate: %s", err)
	}
}

func TestOfcOLEObjectMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcOLEObject()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcOLEObject()
	xml.Unmarshal(buf, v2)
}
