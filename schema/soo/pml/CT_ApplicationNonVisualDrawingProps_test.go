package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_ApplicationNonVisualDrawingPropsConstructor(t *testing.T) {
	v := pml.NewCT_ApplicationNonVisualDrawingProps()
	if v == nil {
		t.Errorf("pml.NewCT_ApplicationNonVisualDrawingProps must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_ApplicationNonVisualDrawingProps should validate: %s", err)
	}
}

func TestCT_ApplicationNonVisualDrawingPropsMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_ApplicationNonVisualDrawingProps()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_ApplicationNonVisualDrawingProps()
	xml.Unmarshal(buf, v2)
}
