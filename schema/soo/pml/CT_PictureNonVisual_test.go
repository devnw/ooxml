package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_PictureNonVisualConstructor(t *testing.T) {
	v := pml.NewCT_PictureNonVisual()
	if v == nil {
		t.Errorf("pml.NewCT_PictureNonVisual must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_PictureNonVisual should validate: %s", err)
	}
}

func TestCT_PictureNonVisualMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_PictureNonVisual()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_PictureNonVisual()
	xml.Unmarshal(buf, v2)
}
