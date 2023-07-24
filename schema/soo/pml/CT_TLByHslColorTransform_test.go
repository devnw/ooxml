package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TLByHslColorTransformConstructor(t *testing.T) {
	v := pml.NewCT_TLByHslColorTransform()
	if v == nil {
		t.Errorf("pml.NewCT_TLByHslColorTransform must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLByHslColorTransform should validate: %s", err)
	}
}

func TestCT_TLByHslColorTransformMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLByHslColorTransform()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLByHslColorTransform()
	xml.Unmarshal(buf, v2)
}
