package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_ColorSchemeAndMappingConstructor(t *testing.T) {
	v := dml.NewCT_ColorSchemeAndMapping()
	if v == nil {
		t.Errorf("dml.NewCT_ColorSchemeAndMapping must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_ColorSchemeAndMapping should validate: %s", err)
	}
}

func TestCT_ColorSchemeAndMappingMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_ColorSchemeAndMapping()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_ColorSchemeAndMapping()
	xml.Unmarshal(buf, v2)
}
