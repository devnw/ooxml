package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_FontConstructor(t *testing.T) {
	v := wml.NewCT_Font()
	if v == nil {
		t.Errorf("wml.NewCT_Font must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Font should validate: %s", err)
	}
}

func TestCT_FontMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Font()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Font()
	xml.Unmarshal(buf, v2)
}
