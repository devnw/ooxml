package spreadsheetDrawing_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/spreadsheetDrawing"
)

func TestCT_GraphicalObjectFrameConstructor(t *testing.T) {
	v := spreadsheetDrawing.NewCT_GraphicalObjectFrame()
	if v == nil {
		t.Errorf("spreadsheetDrawing.NewCT_GraphicalObjectFrame must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed spreadsheetDrawing.CT_GraphicalObjectFrame should validate: %s", err)
	}
}

func TestCT_GraphicalObjectFrameMarshalUnmarshal(t *testing.T) {
	v := spreadsheetDrawing.NewCT_GraphicalObjectFrame()
	buf, _ := xml.Marshal(v)
	v2 := spreadsheetDrawing.NewCT_GraphicalObjectFrame()
	xml.Unmarshal(buf, v2)
}
