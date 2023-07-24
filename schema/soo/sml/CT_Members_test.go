package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_MembersConstructor(t *testing.T) {
	v := sml.NewCT_Members()
	if v == nil {
		t.Errorf("sml.NewCT_Members must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Members should validate: %s", err)
	}
}

func TestCT_MembersMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Members()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Members()
	xml.Unmarshal(buf, v2)
}
