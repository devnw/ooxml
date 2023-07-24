package chartDrawing_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chartDrawing"
)

func TestCT_GraphicFrameNonVisualConstructor(t *testing.T) {
	v := chartDrawing.NewCT_GraphicFrameNonVisual()
	if v == nil {
		t.Errorf("chartDrawing.NewCT_GraphicFrameNonVisual must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chartDrawing.CT_GraphicFrameNonVisual should validate: %s", err)
	}
}

func TestCT_GraphicFrameNonVisualMarshalUnmarshal(t *testing.T) {
	v := chartDrawing.NewCT_GraphicFrameNonVisual()
	buf, _ := xml.Marshal(v)
	v2 := chartDrawing.NewCT_GraphicFrameNonVisual()
	xml.Unmarshal(buf, v2)
}
