package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_GradientStopConstructor(t *testing.T) {
	v := sml.NewCT_GradientStop()
	if v == nil {
		t.Errorf("sml.NewCT_GradientStop must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_GradientStop should validate: %s", err)
	}
}

func TestCT_GradientStopMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_GradientStop()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_GradientStop()
	xml.Unmarshal(buf, v2)
}
