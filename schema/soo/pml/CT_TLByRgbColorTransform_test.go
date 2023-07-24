package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TLByRgbColorTransformConstructor(t *testing.T) {
	v := pml.NewCT_TLByRgbColorTransform()
	if v == nil {
		t.Errorf("pml.NewCT_TLByRgbColorTransform must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLByRgbColorTransform should validate: %s", err)
	}
}

func TestCT_TLByRgbColorTransformMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLByRgbColorTransform()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLByRgbColorTransform()
	xml.Unmarshal(buf, v2)
}
