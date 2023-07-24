package spreadsheetDrawing_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/spreadsheetDrawing"
)

func TestFromConstructor(t *testing.T) {
	v := spreadsheetDrawing.NewFrom()
	if v == nil {
		t.Errorf("spreadsheetDrawing.NewFrom must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed spreadsheetDrawing.From should validate: %s", err)
	}
}

func TestFromMarshalUnmarshal(t *testing.T) {
	v := spreadsheetDrawing.NewFrom()
	buf, _ := xml.Marshal(v)
	v2 := spreadsheetDrawing.NewFrom()
	xml.Unmarshal(buf, v2)
}
