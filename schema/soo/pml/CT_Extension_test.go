package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_ExtensionConstructor(t *testing.T) {
	v := pml.NewCT_Extension()
	if v == nil {
		t.Errorf("pml.NewCT_Extension must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_Extension should validate: %s", err)
	}
}

func TestCT_ExtensionMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_Extension()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_Extension()
	xml.Unmarshal(buf, v2)
}
