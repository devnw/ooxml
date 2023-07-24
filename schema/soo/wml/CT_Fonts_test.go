package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_FontsConstructor(t *testing.T) {
	v := wml.NewCT_Fonts()
	if v == nil {
		t.Errorf("wml.NewCT_Fonts must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Fonts should validate: %s", err)
	}
}

func TestCT_FontsMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Fonts()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Fonts()
	xml.Unmarshal(buf, v2)
}
