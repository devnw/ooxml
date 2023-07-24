package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_FontNameConstructor(t *testing.T) {
	v := sml.NewCT_FontName()
	if v == nil {
		t.Errorf("sml.NewCT_FontName must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_FontName should validate: %s", err)
	}
}

func TestCT_FontNameMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_FontName()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_FontName()
	xml.Unmarshal(buf, v2)
}
