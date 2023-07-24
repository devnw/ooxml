package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_FramesetChoiceConstructor(t *testing.T) {
	v := wml.NewCT_FramesetChoice()
	if v == nil {
		t.Errorf("wml.NewCT_FramesetChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_FramesetChoice should validate: %s", err)
	}
}

func TestCT_FramesetChoiceMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_FramesetChoice()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_FramesetChoice()
	xml.Unmarshal(buf, v2)
}
