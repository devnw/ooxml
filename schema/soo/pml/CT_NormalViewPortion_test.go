package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_NormalViewPortionConstructor(t *testing.T) {
	v := pml.NewCT_NormalViewPortion()
	if v == nil {
		t.Errorf("pml.NewCT_NormalViewPortion must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_NormalViewPortion should validate: %s", err)
	}
}

func TestCT_NormalViewPortionMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_NormalViewPortion()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_NormalViewPortion()
	xml.Unmarshal(buf, v2)
}
