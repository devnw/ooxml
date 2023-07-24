package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_RgbColorConstructor(t *testing.T) {
	v := sml.NewCT_RgbColor()
	if v == nil {
		t.Errorf("sml.NewCT_RgbColor must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_RgbColor should validate: %s", err)
	}
}

func TestCT_RgbColorMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_RgbColor()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_RgbColor()
	xml.Unmarshal(buf, v2)
}
