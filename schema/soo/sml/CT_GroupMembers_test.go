package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_GroupMembersConstructor(t *testing.T) {
	v := sml.NewCT_GroupMembers()
	if v == nil {
		t.Errorf("sml.NewCT_GroupMembers must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_GroupMembers should validate: %s", err)
	}
}

func TestCT_GroupMembersMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_GroupMembers()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_GroupMembers()
	xml.Unmarshal(buf, v2)
}
