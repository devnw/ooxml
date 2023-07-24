package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_RelConstructor(t *testing.T) {
	v := pml.NewCT_Rel()
	if v == nil {
		t.Errorf("pml.NewCT_Rel must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_Rel should validate: %s", err)
	}
}

func TestCT_RelMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_Rel()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_Rel()
	xml.Unmarshal(buf, v2)
}
