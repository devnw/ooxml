package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestWdInlineConstructor(t *testing.T) {
	v := wml.NewWdInline()
	if v == nil {
		t.Errorf("wml.NewWdInline must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.WdInline should validate: %s", err)
	}
}

func TestWdInlineMarshalUnmarshal(t *testing.T) {
	v := wml.NewWdInline()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewWdInline()
	xml.Unmarshal(buf, v2)
}
