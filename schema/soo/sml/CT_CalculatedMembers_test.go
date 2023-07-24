package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_CalculatedMembersConstructor(t *testing.T) {
	v := sml.NewCT_CalculatedMembers()
	if v == nil {
		t.Errorf("sml.NewCT_CalculatedMembers must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_CalculatedMembers should validate: %s", err)
	}
}

func TestCT_CalculatedMembersMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_CalculatedMembers()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_CalculatedMembers()
	xml.Unmarshal(buf, v2)
}
