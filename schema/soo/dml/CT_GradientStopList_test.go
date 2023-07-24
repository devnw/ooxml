package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_GradientStopListConstructor(t *testing.T) {
	v := dml.NewCT_GradientStopList()
	if v == nil {
		t.Errorf("dml.NewCT_GradientStopList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_GradientStopList should validate: %s", err)
	}
}

func TestCT_GradientStopListMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_GradientStopList()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_GradientStopList()
	xml.Unmarshal(buf, v2)
}
