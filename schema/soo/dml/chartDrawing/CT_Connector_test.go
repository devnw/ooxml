package chartDrawing_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chartDrawing"
)

func TestCT_ConnectorConstructor(t *testing.T) {
	v := chartDrawing.NewCT_Connector()
	if v == nil {
		t.Errorf("chartDrawing.NewCT_Connector must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chartDrawing.CT_Connector should validate: %s", err)
	}
}

func TestCT_ConnectorMarshalUnmarshal(t *testing.T) {
	v := chartDrawing.NewCT_Connector()
	buf, _ := xml.Marshal(v)
	v2 := chartDrawing.NewCT_Connector()
	xml.Unmarshal(buf, v2)
}
