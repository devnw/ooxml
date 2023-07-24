package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestWdCT_WrapNoneConstructor(t *testing.T) {
	v := wml.NewWdCT_WrapNone()
	if v == nil {
		t.Errorf("wml.NewWdCT_WrapNone must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.WdCT_WrapNone should validate: %s", err)
	}
}

func TestWdCT_WrapNoneMarshalUnmarshal(t *testing.T) {
	v := wml.NewWdCT_WrapNone()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewWdCT_WrapNone()
	xml.Unmarshal(buf, v2)
}
