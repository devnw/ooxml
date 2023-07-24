package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_GradientFillConstructor(t *testing.T) {
	v := sml.NewCT_GradientFill()
	if v == nil {
		t.Errorf("sml.NewCT_GradientFill must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_GradientFill should validate: %s", err)
	}
}

func TestCT_GradientFillMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_GradientFill()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_GradientFill()
	xml.Unmarshal(buf, v2)
}
