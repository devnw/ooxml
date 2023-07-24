package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestWdCT_WrapThroughConstructor(t *testing.T) {
	v := wml.NewWdCT_WrapThrough()
	if v == nil {
		t.Errorf("wml.NewWdCT_WrapThrough must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.WdCT_WrapThrough should validate: %s", err)
	}
}

func TestWdCT_WrapThroughMarshalUnmarshal(t *testing.T) {
	v := wml.NewWdCT_WrapThrough()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewWdCT_WrapThrough()
	xml.Unmarshal(buf, v2)
}
