package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_KinsokuConstructor(t *testing.T) {
	v := pml.NewCT_Kinsoku()
	if v == nil {
		t.Errorf("pml.NewCT_Kinsoku must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_Kinsoku should validate: %s", err)
	}
}

func TestCT_KinsokuMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_Kinsoku()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_Kinsoku()
	xml.Unmarshal(buf, v2)
}
