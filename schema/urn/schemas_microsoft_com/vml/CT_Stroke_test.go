package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestCT_StrokeConstructor(t *testing.T) {
	v := vml.NewCT_Stroke()
	if v == nil {
		t.Errorf("vml.NewCT_Stroke must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.CT_Stroke should validate: %s", err)
	}
}

func TestCT_StrokeMarshalUnmarshal(t *testing.T) {
	v := vml.NewCT_Stroke()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewCT_Stroke()
	xml.Unmarshal(buf, v2)
}
