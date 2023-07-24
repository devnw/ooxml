package spreadsheetDrawing_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/spreadsheetDrawing"
)

func TestCT_GraphicalObjectFrameNonVisualConstructor(t *testing.T) {
	v := spreadsheetDrawing.NewCT_GraphicalObjectFrameNonVisual()
	if v == nil {
		t.Errorf("spreadsheetDrawing.NewCT_GraphicalObjectFrameNonVisual must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed spreadsheetDrawing.CT_GraphicalObjectFrameNonVisual should validate: %s", err)
	}
}

func TestCT_GraphicalObjectFrameNonVisualMarshalUnmarshal(t *testing.T) {
	v := spreadsheetDrawing.NewCT_GraphicalObjectFrameNonVisual()
	buf, _ := xml.Marshal(v)
	v2 := spreadsheetDrawing.NewCT_GraphicalObjectFrameNonVisual()
	xml.Unmarshal(buf, v2)
}
