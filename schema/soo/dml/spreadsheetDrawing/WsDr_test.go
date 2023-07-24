package spreadsheetDrawing_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/spreadsheetDrawing"
)

func TestWsDrConstructor(t *testing.T) {
	v := spreadsheetDrawing.NewWsDr()
	if v == nil {
		t.Errorf("spreadsheetDrawing.NewWsDr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed spreadsheetDrawing.WsDr should validate: %s", err)
	}
}

func TestWsDrMarshalUnmarshal(t *testing.T) {
	v := spreadsheetDrawing.NewWsDr()
	buf, _ := xml.Marshal(v)
	v2 := spreadsheetDrawing.NewWsDr()
	xml.Unmarshal(buf, v2)
}
