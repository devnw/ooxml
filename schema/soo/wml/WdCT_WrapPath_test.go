package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestWdCT_WrapPathConstructor(t *testing.T) {
	v := wml.NewWdCT_WrapPath()
	if v == nil {
		t.Errorf("wml.NewWdCT_WrapPath must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.WdCT_WrapPath should validate: %s", err)
	}
}

func TestWdCT_WrapPathMarshalUnmarshal(t *testing.T) {
	v := wml.NewWdCT_WrapPath()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewWdCT_WrapPath()
	xml.Unmarshal(buf, v2)
}
