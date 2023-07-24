package spreadsheetDrawing_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/spreadsheetDrawing"
)

func TestCT_DrawingConstructor(t *testing.T) {
	v := spreadsheetDrawing.NewCT_Drawing()
	if v == nil {
		t.Errorf("spreadsheetDrawing.NewCT_Drawing must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed spreadsheetDrawing.CT_Drawing should validate: %s", err)
	}
}

func TestCT_DrawingMarshalUnmarshal(t *testing.T) {
	v := spreadsheetDrawing.NewCT_Drawing()
	buf, _ := xml.Marshal(v)
	v2 := spreadsheetDrawing.NewCT_Drawing()
	xml.Unmarshal(buf, v2)
}
