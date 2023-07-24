package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_FrameLayoutConstructor(t *testing.T) {
	v := wml.NewCT_FrameLayout()
	if v == nil {
		t.Errorf("wml.NewCT_FrameLayout must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_FrameLayout should validate: %s", err)
	}
}

func TestCT_FrameLayoutMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_FrameLayout()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_FrameLayout()
	xml.Unmarshal(buf, v2)
}
