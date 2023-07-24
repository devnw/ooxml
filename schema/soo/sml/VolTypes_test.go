package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestVolTypesConstructor(t *testing.T) {
	v := sml.NewVolTypes()
	if v == nil {
		t.Errorf("sml.NewVolTypes must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.VolTypes should validate: %s", err)
	}
}

func TestVolTypesMarshalUnmarshal(t *testing.T) {
	v := sml.NewVolTypes()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewVolTypes()
	xml.Unmarshal(buf, v2)
}
