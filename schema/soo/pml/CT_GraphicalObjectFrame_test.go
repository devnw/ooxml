package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_GraphicalObjectFrameConstructor(t *testing.T) {
	v := pml.NewCT_GraphicalObjectFrame()
	if v == nil {
		t.Errorf("pml.NewCT_GraphicalObjectFrame must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_GraphicalObjectFrame should validate: %s", err)
	}
}

func TestCT_GraphicalObjectFrameMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_GraphicalObjectFrame()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_GraphicalObjectFrame()
	xml.Unmarshal(buf, v2)
}
