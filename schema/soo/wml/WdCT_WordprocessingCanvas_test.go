package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestWdCT_WordprocessingCanvasConstructor(t *testing.T) {
	v := wml.NewWdCT_WordprocessingCanvas()
	if v == nil {
		t.Errorf("wml.NewWdCT_WordprocessingCanvas must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.WdCT_WordprocessingCanvas should validate: %s", err)
	}
}

func TestWdCT_WordprocessingCanvasMarshalUnmarshal(t *testing.T) {
	v := wml.NewWdCT_WordprocessingCanvas()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewWdCT_WordprocessingCanvas()
	xml.Unmarshal(buf, v2)
}
