package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_UnderlineConstructor(t *testing.T) {
	v := wml.NewCT_Underline()
	if v == nil {
		t.Errorf("wml.NewCT_Underline must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Underline should validate: %s", err)
	}
}

func TestCT_UnderlineMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Underline()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Underline()
	xml.Unmarshal(buf, v2)
}
