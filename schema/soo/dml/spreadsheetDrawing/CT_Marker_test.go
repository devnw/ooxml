package spreadsheetDrawing_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/spreadsheetDrawing"
)

func TestCT_MarkerConstructor(t *testing.T) {
	v := spreadsheetDrawing.NewCT_Marker()
	if v == nil {
		t.Errorf("spreadsheetDrawing.NewCT_Marker must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed spreadsheetDrawing.CT_Marker should validate: %s", err)
	}
}

func TestCT_MarkerMarshalUnmarshal(t *testing.T) {
	v := spreadsheetDrawing.NewCT_Marker()
	buf, _ := xml.Marshal(v)
	v2 := spreadsheetDrawing.NewCT_Marker()
	xml.Unmarshal(buf, v2)
}
