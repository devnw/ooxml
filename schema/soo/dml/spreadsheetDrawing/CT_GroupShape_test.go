package spreadsheetDrawing_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/spreadsheetDrawing"
)

func TestCT_GroupShapeConstructor(t *testing.T) {
	v := spreadsheetDrawing.NewCT_GroupShape()
	if v == nil {
		t.Errorf("spreadsheetDrawing.NewCT_GroupShape must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed spreadsheetDrawing.CT_GroupShape should validate: %s", err)
	}
}

func TestCT_GroupShapeMarshalUnmarshal(t *testing.T) {
	v := spreadsheetDrawing.NewCT_GroupShape()
	buf, _ := xml.Marshal(v)
	v2 := spreadsheetDrawing.NewCT_GroupShape()
	xml.Unmarshal(buf, v2)
}
