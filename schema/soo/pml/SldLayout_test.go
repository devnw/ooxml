package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestSldLayoutConstructor(t *testing.T) {
	v := pml.NewSldLayout()
	if v == nil {
		t.Errorf("pml.NewSldLayout must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.SldLayout should validate: %s", err)
	}
}

func TestSldLayoutMarshalUnmarshal(t *testing.T) {
	v := pml.NewSldLayout()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewSldLayout()
	xml.Unmarshal(buf, v2)
}
