package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestCT_CurveConstructor(t *testing.T) {
	v := vml.NewCT_Curve()
	if v == nil {
		t.Errorf("vml.NewCT_Curve must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.CT_Curve should validate: %s", err)
	}
}

func TestCT_CurveMarshalUnmarshal(t *testing.T) {
	v := vml.NewCT_Curve()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewCT_Curve()
	xml.Unmarshal(buf, v2)
}
