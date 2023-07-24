package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_FontsConstructor(t *testing.T) {
	v := sml.NewCT_Fonts()
	if v == nil {
		t.Errorf("sml.NewCT_Fonts must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Fonts should validate: %s", err)
	}
}

func TestCT_FontsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Fonts()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Fonts()
	xml.Unmarshal(buf, v2)
}
