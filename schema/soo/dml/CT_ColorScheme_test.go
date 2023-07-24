package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_ColorSchemeConstructor(t *testing.T) {
	v := dml.NewCT_ColorScheme()
	if v == nil {
		t.Errorf("dml.NewCT_ColorScheme must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_ColorScheme should validate: %s", err)
	}
}

func TestCT_ColorSchemeMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_ColorScheme()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_ColorScheme()
	xml.Unmarshal(buf, v2)
}
