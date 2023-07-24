package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestRectConstructor(t *testing.T) {
	v := vml.NewRect()
	if v == nil {
		t.Errorf("vml.NewRect must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.Rect should validate: %s", err)
	}
}

func TestRectMarshalUnmarshal(t *testing.T) {
	v := vml.NewRect()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewRect()
	xml.Unmarshal(buf, v2)
}
