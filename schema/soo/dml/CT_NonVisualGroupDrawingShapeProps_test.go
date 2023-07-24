package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_NonVisualGroupDrawingShapePropsConstructor(t *testing.T) {
	v := dml.NewCT_NonVisualGroupDrawingShapeProps()
	if v == nil {
		t.Errorf("dml.NewCT_NonVisualGroupDrawingShapeProps must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_NonVisualGroupDrawingShapeProps should validate: %s", err)
	}
}

func TestCT_NonVisualGroupDrawingShapePropsMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_NonVisualGroupDrawingShapeProps()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_NonVisualGroupDrawingShapeProps()
	xml.Unmarshal(buf, v2)
}
