package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_FramesetConstructor(t *testing.T) {
	v := wml.NewCT_Frameset()
	if v == nil {
		t.Errorf("wml.NewCT_Frameset must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Frameset should validate: %s", err)
	}
}

func TestCT_FramesetMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Frameset()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Frameset()
	xml.Unmarshal(buf, v2)
}
