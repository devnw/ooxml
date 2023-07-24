package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_FontSchemeConstructor(t *testing.T) {
	v := sml.NewCT_FontScheme()
	if v == nil {
		t.Errorf("sml.NewCT_FontScheme must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_FontScheme should validate: %s", err)
	}
}

func TestCT_FontSchemeMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_FontScheme()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_FontScheme()
	xml.Unmarshal(buf, v2)
}
