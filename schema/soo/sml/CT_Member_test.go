package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_MemberConstructor(t *testing.T) {
	v := sml.NewCT_Member()
	if v == nil {
		t.Errorf("sml.NewCT_Member must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Member should validate: %s", err)
	}
}

func TestCT_MemberMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Member()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Member()
	xml.Unmarshal(buf, v2)
}
