package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_GroupsConstructor(t *testing.T) {
	v := sml.NewCT_Groups()
	if v == nil {
		t.Errorf("sml.NewCT_Groups must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Groups should validate: %s", err)
	}
}

func TestCT_GroupsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Groups()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Groups()
	xml.Unmarshal(buf, v2)
}
