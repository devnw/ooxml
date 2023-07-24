package spreadsheetDrawing_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/spreadsheetDrawing"
)

func TestCT_OneCellAnchorConstructor(t *testing.T) {
	v := spreadsheetDrawing.NewCT_OneCellAnchor()
	if v == nil {
		t.Errorf("spreadsheetDrawing.NewCT_OneCellAnchor must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed spreadsheetDrawing.CT_OneCellAnchor should validate: %s", err)
	}
}

func TestCT_OneCellAnchorMarshalUnmarshal(t *testing.T) {
	v := spreadsheetDrawing.NewCT_OneCellAnchor()
	buf, _ := xml.Marshal(v)
	v2 := spreadsheetDrawing.NewCT_OneCellAnchor()
	xml.Unmarshal(buf, v2)
}
