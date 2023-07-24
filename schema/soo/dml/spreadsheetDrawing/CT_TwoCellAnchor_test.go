package spreadsheetDrawing_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/spreadsheetDrawing"
)

func TestCT_TwoCellAnchorConstructor(t *testing.T) {
	v := spreadsheetDrawing.NewCT_TwoCellAnchor()
	if v == nil {
		t.Errorf("spreadsheetDrawing.NewCT_TwoCellAnchor must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed spreadsheetDrawing.CT_TwoCellAnchor should validate: %s", err)
	}
}

func TestCT_TwoCellAnchorMarshalUnmarshal(t *testing.T) {
	v := spreadsheetDrawing.NewCT_TwoCellAnchor()
	buf, _ := xml.Marshal(v)
	v2 := spreadsheetDrawing.NewCT_TwoCellAnchor()
	xml.Unmarshal(buf, v2)
}
