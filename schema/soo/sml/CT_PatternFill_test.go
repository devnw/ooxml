package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_PatternFillConstructor(t *testing.T) {
	v := sml.NewCT_PatternFill()
	if v == nil {
		t.Errorf("sml.NewCT_PatternFill must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_PatternFill should validate: %s", err)
	}
}

func TestCT_PatternFillMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_PatternFill()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_PatternFill()
	xml.Unmarshal(buf, v2)
}
