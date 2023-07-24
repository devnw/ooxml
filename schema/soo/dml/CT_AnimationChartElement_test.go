package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_AnimationChartElementConstructor(t *testing.T) {
	v := dml.NewCT_AnimationChartElement()
	if v == nil {
		t.Errorf("dml.NewCT_AnimationChartElement must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_AnimationChartElement should validate: %s", err)
	}
}

func TestCT_AnimationChartElementMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_AnimationChartElement()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_AnimationChartElement()
	xml.Unmarshal(buf, v2)
}
