package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_GvmlUseShapeRectangleConstructor(t *testing.T) {
	v := dml.NewCT_GvmlUseShapeRectangle()
	if v == nil {
		t.Errorf("dml.NewCT_GvmlUseShapeRectangle must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_GvmlUseShapeRectangle should validate: %s", err)
	}
}

func TestCT_GvmlUseShapeRectangleMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_GvmlUseShapeRectangle()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_GvmlUseShapeRectangle()
	xml.Unmarshal(buf, v2)
}
