package spreadsheetDrawing_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/spreadsheetDrawing"
)

func TestEG_ObjectChoicesChoiceConstructor(t *testing.T) {
	v := spreadsheetDrawing.NewEG_ObjectChoicesChoice()
	if v == nil {
		t.Errorf("spreadsheetDrawing.NewEG_ObjectChoicesChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed spreadsheetDrawing.EG_ObjectChoicesChoice should validate: %s", err)
	}
}

func TestEG_ObjectChoicesChoiceMarshalUnmarshal(t *testing.T) {
	v := spreadsheetDrawing.NewEG_ObjectChoicesChoice()
	buf, _ := xml.Marshal(v)
	v2 := spreadsheetDrawing.NewEG_ObjectChoicesChoice()
	xml.Unmarshal(buf, v2)
}
