package spreadsheetDrawing_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/spreadsheetDrawing"
)

func TestEG_ObjectChoicesConstructor(t *testing.T) {
	v := spreadsheetDrawing.NewEG_ObjectChoices()
	if v == nil {
		t.Errorf("spreadsheetDrawing.NewEG_ObjectChoices must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed spreadsheetDrawing.EG_ObjectChoices should validate: %s", err)
	}
}

func TestEG_ObjectChoicesMarshalUnmarshal(t *testing.T) {
	v := spreadsheetDrawing.NewEG_ObjectChoices()
	buf, _ := xml.Marshal(v)
	v2 := spreadsheetDrawing.NewEG_ObjectChoices()
	xml.Unmarshal(buf, v2)
}
