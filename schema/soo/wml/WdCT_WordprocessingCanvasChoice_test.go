package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestWdCT_WordprocessingCanvasChoiceConstructor(t *testing.T) {
	v := wml.NewWdCT_WordprocessingCanvasChoice()
	if v == nil {
		t.Errorf("wml.NewWdCT_WordprocessingCanvasChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.WdCT_WordprocessingCanvasChoice should validate: %s", err)
	}
}

func TestWdCT_WordprocessingCanvasChoiceMarshalUnmarshal(t *testing.T) {
	v := wml.NewWdCT_WordprocessingCanvasChoice()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewWdCT_WordprocessingCanvasChoice()
	xml.Unmarshal(buf, v2)
}
