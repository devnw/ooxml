package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestWdCT_GraphicFrameConstructor(t *testing.T) {
	v := wml.NewWdCT_GraphicFrame()
	if v == nil {
		t.Errorf("wml.NewWdCT_GraphicFrame must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.WdCT_GraphicFrame should validate: %s", err)
	}
}

func TestWdCT_GraphicFrameMarshalUnmarshal(t *testing.T) {
	v := wml.NewWdCT_GraphicFrame()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewWdCT_GraphicFrame()
	xml.Unmarshal(buf, v2)
}
