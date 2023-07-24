package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_FrameScrollbarConstructor(t *testing.T) {
	v := wml.NewCT_FrameScrollbar()
	if v == nil {
		t.Errorf("wml.NewCT_FrameScrollbar must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_FrameScrollbar should validate: %s", err)
	}
}

func TestCT_FrameScrollbarMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_FrameScrollbar()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_FrameScrollbar()
	xml.Unmarshal(buf, v2)
}
