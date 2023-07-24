package spreadsheetDrawing_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/spreadsheetDrawing"
)

func TestCT_AnchorClientDataConstructor(t *testing.T) {
	v := spreadsheetDrawing.NewCT_AnchorClientData()
	if v == nil {
		t.Errorf("spreadsheetDrawing.NewCT_AnchorClientData must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed spreadsheetDrawing.CT_AnchorClientData should validate: %s", err)
	}
}

func TestCT_AnchorClientDataMarshalUnmarshal(t *testing.T) {
	v := spreadsheetDrawing.NewCT_AnchorClientData()
	buf, _ := xml.Marshal(v)
	v2 := spreadsheetDrawing.NewCT_AnchorClientData()
	xml.Unmarshal(buf, v2)
}
