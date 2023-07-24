package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_GraphicalObjectFrameNonVisualConstructor(t *testing.T) {
	v := pml.NewCT_GraphicalObjectFrameNonVisual()
	if v == nil {
		t.Errorf("pml.NewCT_GraphicalObjectFrameNonVisual must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_GraphicalObjectFrameNonVisual should validate: %s", err)
	}
}

func TestCT_GraphicalObjectFrameNonVisualMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_GraphicalObjectFrameNonVisual()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_GraphicalObjectFrameNonVisual()
	xml.Unmarshal(buf, v2)
}
