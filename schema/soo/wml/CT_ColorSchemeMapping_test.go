package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_ColorSchemeMappingConstructor(t *testing.T) {
	v := wml.NewCT_ColorSchemeMapping()
	if v == nil {
		t.Errorf("wml.NewCT_ColorSchemeMapping must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_ColorSchemeMapping should validate: %s", err)
	}
}

func TestCT_ColorSchemeMappingMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_ColorSchemeMapping()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_ColorSchemeMapping()
	xml.Unmarshal(buf, v2)
}
