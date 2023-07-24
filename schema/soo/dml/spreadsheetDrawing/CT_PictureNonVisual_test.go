package spreadsheetDrawing_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/spreadsheetDrawing"
)

func TestCT_PictureNonVisualConstructor(t *testing.T) {
	v := spreadsheetDrawing.NewCT_PictureNonVisual()
	if v == nil {
		t.Errorf("spreadsheetDrawing.NewCT_PictureNonVisual must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed spreadsheetDrawing.CT_PictureNonVisual should validate: %s", err)
	}
}

func TestCT_PictureNonVisualMarshalUnmarshal(t *testing.T) {
	v := spreadsheetDrawing.NewCT_PictureNonVisual()
	buf, _ := xml.Marshal(v)
	v2 := spreadsheetDrawing.NewCT_PictureNonVisual()
	xml.Unmarshal(buf, v2)
}
