package spreadsheetDrawing_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/spreadsheetDrawing"
)

func TestCT_RelConstructor(t *testing.T) {
	v := spreadsheetDrawing.NewCT_Rel()
	if v == nil {
		t.Errorf("spreadsheetDrawing.NewCT_Rel must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed spreadsheetDrawing.CT_Rel should validate: %s", err)
	}
}

func TestCT_RelMarshalUnmarshal(t *testing.T) {
	v := spreadsheetDrawing.NewCT_Rel()
	buf, _ := xml.Marshal(v)
	v2 := spreadsheetDrawing.NewCT_Rel()
	xml.Unmarshal(buf, v2)
}
