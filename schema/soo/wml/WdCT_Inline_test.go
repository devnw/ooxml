package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestWdCT_InlineConstructor(t *testing.T) {
	v := wml.NewWdCT_Inline()
	if v == nil {
		t.Errorf("wml.NewWdCT_Inline must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.WdCT_Inline should validate: %s", err)
	}
}

func TestWdCT_InlineMarshalUnmarshal(t *testing.T) {
	v := wml.NewWdCT_Inline()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewWdCT_Inline()
	xml.Unmarshal(buf, v2)
}
