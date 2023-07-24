package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_FramesetSplitbarConstructor(t *testing.T) {
	v := wml.NewCT_FramesetSplitbar()
	if v == nil {
		t.Errorf("wml.NewCT_FramesetSplitbar must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_FramesetSplitbar should validate: %s", err)
	}
}

func TestCT_FramesetSplitbarMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_FramesetSplitbar()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_FramesetSplitbar()
	xml.Unmarshal(buf, v2)
}
