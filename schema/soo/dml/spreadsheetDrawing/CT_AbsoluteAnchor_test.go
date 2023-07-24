package spreadsheetDrawing_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/spreadsheetDrawing"
)

func TestCT_AbsoluteAnchorConstructor(t *testing.T) {
	v := spreadsheetDrawing.NewCT_AbsoluteAnchor()
	if v == nil {
		t.Errorf("spreadsheetDrawing.NewCT_AbsoluteAnchor must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed spreadsheetDrawing.CT_AbsoluteAnchor should validate: %s", err)
	}
}

func TestCT_AbsoluteAnchorMarshalUnmarshal(t *testing.T) {
	v := spreadsheetDrawing.NewCT_AbsoluteAnchor()
	buf, _ := xml.Marshal(v)
	v2 := spreadsheetDrawing.NewCT_AbsoluteAnchor()
	xml.Unmarshal(buf, v2)
}
