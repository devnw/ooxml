package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestEG_ShapeElementsConstructor(t *testing.T) {
	v := vml.NewEG_ShapeElements()
	if v == nil {
		t.Errorf("vml.NewEG_ShapeElements must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.EG_ShapeElements should validate: %s", err)
	}
}

func TestEG_ShapeElementsMarshalUnmarshal(t *testing.T) {
	v := vml.NewEG_ShapeElements()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewEG_ShapeElements()
	xml.Unmarshal(buf, v2)
}
