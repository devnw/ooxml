package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_PlaceholderConstructor(t *testing.T) {
	v := pml.NewCT_Placeholder()
	if v == nil {
		t.Errorf("pml.NewCT_Placeholder must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_Placeholder should validate: %s", err)
	}
}

func TestCT_PlaceholderMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_Placeholder()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_Placeholder()
	xml.Unmarshal(buf, v2)
}
