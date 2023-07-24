package spreadsheetDrawing_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/spreadsheetDrawing"
)

func TestCT_ConnectorConstructor(t *testing.T) {
	v := spreadsheetDrawing.NewCT_Connector()
	if v == nil {
		t.Errorf("spreadsheetDrawing.NewCT_Connector must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed spreadsheetDrawing.CT_Connector should validate: %s", err)
	}
}

func TestCT_ConnectorMarshalUnmarshal(t *testing.T) {
	v := spreadsheetDrawing.NewCT_Connector()
	buf, _ := xml.Marshal(v)
	v2 := spreadsheetDrawing.NewCT_Connector()
	xml.Unmarshal(buf, v2)
}
