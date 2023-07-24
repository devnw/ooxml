package chartDrawing_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chartDrawing"
)

func TestCT_GroupShapeNonVisualConstructor(t *testing.T) {
	v := chartDrawing.NewCT_GroupShapeNonVisual()
	if v == nil {
		t.Errorf("chartDrawing.NewCT_GroupShapeNonVisual must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chartDrawing.CT_GroupShapeNonVisual should validate: %s", err)
	}
}

func TestCT_GroupShapeNonVisualMarshalUnmarshal(t *testing.T) {
	v := chartDrawing.NewCT_GroupShapeNonVisual()
	buf, _ := xml.Marshal(v)
	v2 := chartDrawing.NewCT_GroupShapeNonVisual()
	xml.Unmarshal(buf, v2)
}
