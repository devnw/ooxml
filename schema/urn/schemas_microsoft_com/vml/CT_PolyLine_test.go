package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestCT_PolyLineConstructor(t *testing.T) {
	v := vml.NewCT_PolyLine()
	if v == nil {
		t.Errorf("vml.NewCT_PolyLine must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.CT_PolyLine should validate: %s", err)
	}
}

func TestCT_PolyLineMarshalUnmarshal(t *testing.T) {
	v := vml.NewCT_PolyLine()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewCT_PolyLine()
	xml.Unmarshal(buf, v2)
}
