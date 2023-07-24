package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_EmptyConstructor(t *testing.T) {
	v := pml.NewCT_Empty()
	if v == nil {
		t.Errorf("pml.NewCT_Empty must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_Empty should validate: %s", err)
	}
}

func TestCT_EmptyMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_Empty()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_Empty()
	xml.Unmarshal(buf, v2)
}
