package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_FramePrConstructor(t *testing.T) {
	v := wml.NewCT_FramePr()
	if v == nil {
		t.Errorf("wml.NewCT_FramePr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_FramePr should validate: %s", err)
	}
}

func TestCT_FramePrMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_FramePr()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_FramePr()
	xml.Unmarshal(buf, v2)
}
