package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_MemberPropertiesConstructor(t *testing.T) {
	v := sml.NewCT_MemberProperties()
	if v == nil {
		t.Errorf("sml.NewCT_MemberProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_MemberProperties should validate: %s", err)
	}
}

func TestCT_MemberPropertiesMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_MemberProperties()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_MemberProperties()
	xml.Unmarshal(buf, v2)
}
