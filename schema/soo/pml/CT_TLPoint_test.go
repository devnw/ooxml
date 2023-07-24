package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TLPointConstructor(t *testing.T) {
	v := pml.NewCT_TLPoint()
	if v == nil {
		t.Errorf("pml.NewCT_TLPoint must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLPoint should validate: %s", err)
	}
}

func TestCT_TLPointMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLPoint()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLPoint()
	xml.Unmarshal(buf, v2)
}
