package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestArcConstructor(t *testing.T) {
	v := vml.NewArc()
	if v == nil {
		t.Errorf("vml.NewArc must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.Arc should validate: %s", err)
	}
}

func TestArcMarshalUnmarshal(t *testing.T) {
	v := vml.NewArc()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewArc()
	xml.Unmarshal(buf, v2)
}
