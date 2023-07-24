package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestSldConstructor(t *testing.T) {
	v := pml.NewSld()
	if v == nil {
		t.Errorf("pml.NewSld must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.Sld should validate: %s", err)
	}
}

func TestSldMarshalUnmarshal(t *testing.T) {
	v := pml.NewSld()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewSld()
	xml.Unmarshal(buf, v2)
}
