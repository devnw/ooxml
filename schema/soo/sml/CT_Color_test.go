package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ColorConstructor(t *testing.T) {
	v := sml.NewCT_Color()
	if v == nil {
		t.Errorf("sml.NewCT_Color must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Color should validate: %s", err)
	}
}

func TestCT_ColorMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Color()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Color()
	xml.Unmarshal(buf, v2)
}
