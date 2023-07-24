package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_GroupMemberConstructor(t *testing.T) {
	v := sml.NewCT_GroupMember()
	if v == nil {
		t.Errorf("sml.NewCT_GroupMember must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_GroupMember should validate: %s", err)
	}
}

func TestCT_GroupMemberMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_GroupMember()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_GroupMember()
	xml.Unmarshal(buf, v2)
}
