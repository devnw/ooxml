package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_NonVisualDrawingShapePropsConstructor(t *testing.T) {
	v := dml.NewCT_NonVisualDrawingShapeProps()
	if v == nil {
		t.Errorf("dml.NewCT_NonVisualDrawingShapeProps must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_NonVisualDrawingShapeProps should validate: %s", err)
	}
}

func TestCT_NonVisualDrawingShapePropsMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_NonVisualDrawingShapeProps()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_NonVisualDrawingShapeProps()
	xml.Unmarshal(buf, v2)
}
