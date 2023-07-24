package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_SlideTimingConstructor(t *testing.T) {
	v := pml.NewCT_SlideTiming()
	if v == nil {
		t.Errorf("pml.NewCT_SlideTiming must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_SlideTiming should validate: %s", err)
	}
}

func TestCT_SlideTimingMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_SlideTiming()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_SlideTiming()
	xml.Unmarshal(buf, v2)
}
