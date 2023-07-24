package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_EastAsianLayoutConstructor(t *testing.T) {
	v := wml.NewCT_EastAsianLayout()
	if v == nil {
		t.Errorf("wml.NewCT_EastAsianLayout must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_EastAsianLayout should validate: %s", err)
	}
}

func TestCT_EastAsianLayoutMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_EastAsianLayout()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_EastAsianLayout()
	xml.Unmarshal(buf, v2)
}
