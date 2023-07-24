package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_HyperlinkConstructor(t *testing.T) {
	v := sml.NewCT_Hyperlink()
	if v == nil {
		t.Errorf("sml.NewCT_Hyperlink must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Hyperlink should validate: %s", err)
	}
}

func TestCT_HyperlinkMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Hyperlink()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Hyperlink()
	xml.Unmarshal(buf, v2)
}
