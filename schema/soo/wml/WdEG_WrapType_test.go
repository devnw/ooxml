package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestWdEG_WrapTypeConstructor(t *testing.T) {
	v := wml.NewWdEG_WrapType()
	if v == nil {
		t.Errorf("wml.NewWdEG_WrapType must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.WdEG_WrapType should validate: %s", err)
	}
}

func TestWdEG_WrapTypeMarshalUnmarshal(t *testing.T) {
	v := wml.NewWdEG_WrapType()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewWdEG_WrapType()
	xml.Unmarshal(buf, v2)
}
