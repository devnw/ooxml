package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_VolTypeConstructor(t *testing.T) {
	v := sml.NewCT_VolType()
	if v == nil {
		t.Errorf("sml.NewCT_VolType must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_VolType should validate: %s", err)
	}
}

func TestCT_VolTypeMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_VolType()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_VolType()
	xml.Unmarshal(buf, v2)
}
