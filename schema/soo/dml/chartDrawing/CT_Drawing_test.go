package chartDrawing_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chartDrawing"
)

func TestCT_DrawingConstructor(t *testing.T) {
	v := chartDrawing.NewCT_Drawing()
	if v == nil {
		t.Errorf("chartDrawing.NewCT_Drawing must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chartDrawing.CT_Drawing should validate: %s", err)
	}
}

func TestCT_DrawingMarshalUnmarshal(t *testing.T) {
	v := chartDrawing.NewCT_Drawing()
	buf, _ := xml.Marshal(v)
	v2 := chartDrawing.NewCT_Drawing()
	xml.Unmarshal(buf, v2)
}
