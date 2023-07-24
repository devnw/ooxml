package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_HyperlinkConstructor(t *testing.T) {
	v := wml.NewCT_Hyperlink()
	if v == nil {
		t.Errorf("wml.NewCT_Hyperlink must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Hyperlink should validate: %s", err)
	}
}

func TestCT_HyperlinkMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Hyperlink()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Hyperlink()
	xml.Unmarshal(buf, v2)
}
