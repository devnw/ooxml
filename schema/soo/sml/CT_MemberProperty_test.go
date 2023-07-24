package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_MemberPropertyConstructor(t *testing.T) {
	v := sml.NewCT_MemberProperty()
	if v == nil {
		t.Errorf("sml.NewCT_MemberProperty must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_MemberProperty should validate: %s", err)
	}
}

func TestCT_MemberPropertyMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_MemberProperty()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_MemberProperty()
	xml.Unmarshal(buf, v2)
}
