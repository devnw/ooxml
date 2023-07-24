package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestStrokeConstructor(t *testing.T) {
	v := vml.NewStroke()
	if v == nil {
		t.Errorf("vml.NewStroke must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.Stroke should validate: %s", err)
	}
}

func TestStrokeMarshalUnmarshal(t *testing.T) {
	v := vml.NewStroke()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewStroke()
	xml.Unmarshal(buf, v2)
}
