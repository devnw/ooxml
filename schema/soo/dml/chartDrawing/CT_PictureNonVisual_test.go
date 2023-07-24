package chartDrawing_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chartDrawing"
)

func TestCT_PictureNonVisualConstructor(t *testing.T) {
	v := chartDrawing.NewCT_PictureNonVisual()
	if v == nil {
		t.Errorf("chartDrawing.NewCT_PictureNonVisual must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chartDrawing.CT_PictureNonVisual should validate: %s", err)
	}
}

func TestCT_PictureNonVisualMarshalUnmarshal(t *testing.T) {
	v := chartDrawing.NewCT_PictureNonVisual()
	buf, _ := xml.Marshal(v)
	v2 := chartDrawing.NewCT_PictureNonVisual()
	xml.Unmarshal(buf, v2)
}
