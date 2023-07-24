package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcCT_StrokeChildConstructor(t *testing.T) {
	v := vml.NewOfcCT_StrokeChild()
	if v == nil {
		t.Errorf("vml.NewOfcCT_StrokeChild must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcCT_StrokeChild should validate: %s", err)
	}
}

func TestOfcCT_StrokeChildMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcCT_StrokeChild()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcCT_StrokeChild()
	xml.Unmarshal(buf, v2)
}
