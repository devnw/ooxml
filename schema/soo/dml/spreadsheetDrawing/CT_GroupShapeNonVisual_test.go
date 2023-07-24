package spreadsheetDrawing_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/spreadsheetDrawing"
)

func TestCT_GroupShapeNonVisualConstructor(t *testing.T) {
	v := spreadsheetDrawing.NewCT_GroupShapeNonVisual()
	if v == nil {
		t.Errorf("spreadsheetDrawing.NewCT_GroupShapeNonVisual must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed spreadsheetDrawing.CT_GroupShapeNonVisual should validate: %s", err)
	}
}

func TestCT_GroupShapeNonVisualMarshalUnmarshal(t *testing.T) {
	v := spreadsheetDrawing.NewCT_GroupShapeNonVisual()
	buf, _ := xml.Marshal(v)
	v2 := spreadsheetDrawing.NewCT_GroupShapeNonVisual()
	xml.Unmarshal(buf, v2)
}
