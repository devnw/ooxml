package spreadsheetDrawing_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/spreadsheetDrawing"
)

func TestToConstructor(t *testing.T) {
	v := spreadsheetDrawing.NewTo()
	if v == nil {
		t.Errorf("spreadsheetDrawing.NewTo must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed spreadsheetDrawing.To should validate: %s", err)
	}
}

func TestToMarshalUnmarshal(t *testing.T) {
	v := spreadsheetDrawing.NewTo()
	buf, _ := xml.Marshal(v)
	v2 := spreadsheetDrawing.NewTo()
	xml.Unmarshal(buf, v2)
}
