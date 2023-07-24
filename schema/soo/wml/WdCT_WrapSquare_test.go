package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestWdCT_WrapSquareConstructor(t *testing.T) {
	v := wml.NewWdCT_WrapSquare()
	if v == nil {
		t.Errorf("wml.NewWdCT_WrapSquare must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.WdCT_WrapSquare should validate: %s", err)
	}
}

func TestWdCT_WrapSquareMarshalUnmarshal(t *testing.T) {
	v := wml.NewWdCT_WrapSquare()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewWdCT_WrapSquare()
	xml.Unmarshal(buf, v2)
}
