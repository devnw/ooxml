package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_CalculatedMemberConstructor(t *testing.T) {
	v := sml.NewCT_CalculatedMember()
	if v == nil {
		t.Errorf("sml.NewCT_CalculatedMember must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_CalculatedMember should validate: %s", err)
	}
}

func TestCT_CalculatedMemberMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_CalculatedMember()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_CalculatedMember()
	xml.Unmarshal(buf, v2)
}
