package chartDrawing_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chartDrawing"
)

func TestCT_ShapeNonVisualConstructor(t *testing.T) {
	v := chartDrawing.NewCT_ShapeNonVisual()
	if v == nil {
		t.Errorf("chartDrawing.NewCT_ShapeNonVisual must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chartDrawing.CT_ShapeNonVisual should validate: %s", err)
	}
}

func TestCT_ShapeNonVisualMarshalUnmarshal(t *testing.T) {
	v := chartDrawing.NewCT_ShapeNonVisual()
	buf, _ := xml.Marshal(v)
	v2 := chartDrawing.NewCT_ShapeNonVisual()
	xml.Unmarshal(buf, v2)
}
