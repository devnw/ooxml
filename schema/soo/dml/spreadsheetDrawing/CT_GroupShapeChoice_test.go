package spreadsheetDrawing_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/spreadsheetDrawing"
)

func TestCT_GroupShapeChoiceConstructor(t *testing.T) {
	v := spreadsheetDrawing.NewCT_GroupShapeChoice()
	if v == nil {
		t.Errorf("spreadsheetDrawing.NewCT_GroupShapeChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed spreadsheetDrawing.CT_GroupShapeChoice should validate: %s", err)
	}
}

func TestCT_GroupShapeChoiceMarshalUnmarshal(t *testing.T) {
	v := spreadsheetDrawing.NewCT_GroupShapeChoice()
	buf, _ := xml.Marshal(v)
	v2 := spreadsheetDrawing.NewCT_GroupShapeChoice()
	xml.Unmarshal(buf, v2)
}
