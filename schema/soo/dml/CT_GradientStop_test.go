package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_GradientStopConstructor(t *testing.T) {
	v := dml.NewCT_GradientStop()
	if v == nil {
		t.Errorf("dml.NewCT_GradientStop must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_GradientStop should validate: %s", err)
	}
}

func TestCT_GradientStopMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_GradientStop()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_GradientStop()
	xml.Unmarshal(buf, v2)
}
