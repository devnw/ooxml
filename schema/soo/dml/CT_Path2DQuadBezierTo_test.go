package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_Path2DQuadBezierToConstructor(t *testing.T) {
	v := dml.NewCT_Path2DQuadBezierTo()
	if v == nil {
		t.Errorf("dml.NewCT_Path2DQuadBezierTo must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_Path2DQuadBezierTo should validate: %s", err)
	}
}

func TestCT_Path2DQuadBezierToMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_Path2DQuadBezierTo()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_Path2DQuadBezierTo()
	xml.Unmarshal(buf, v2)
}
