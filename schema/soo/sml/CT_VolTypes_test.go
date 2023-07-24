package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_VolTypesConstructor(t *testing.T) {
	v := sml.NewCT_VolTypes()
	if v == nil {
		t.Errorf("sml.NewCT_VolTypes must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_VolTypes should validate: %s", err)
	}
}

func TestCT_VolTypesMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_VolTypes()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_VolTypes()
	xml.Unmarshal(buf, v2)
}
