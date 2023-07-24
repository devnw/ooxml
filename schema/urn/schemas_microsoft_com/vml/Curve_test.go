package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestCurveConstructor(t *testing.T) {
	v := vml.NewCurve()
	if v == nil {
		t.Errorf("vml.NewCurve must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.Curve should validate: %s", err)
	}
}

func TestCurveMarshalUnmarshal(t *testing.T) {
	v := vml.NewCurve()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewCurve()
	xml.Unmarshal(buf, v2)
}
