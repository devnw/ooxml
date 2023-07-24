package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_FontFamilyConstructor(t *testing.T) {
	v := wml.NewCT_FontFamily()
	if v == nil {
		t.Errorf("wml.NewCT_FontFamily must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_FontFamily should validate: %s", err)
	}
}

func TestCT_FontFamilyMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_FontFamily()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_FontFamily()
	xml.Unmarshal(buf, v2)
}
