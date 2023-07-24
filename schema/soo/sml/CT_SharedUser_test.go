package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_SharedUserConstructor(t *testing.T) {
	v := sml.NewCT_SharedUser()
	if v == nil {
		t.Errorf("sml.NewCT_SharedUser must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_SharedUser should validate: %s", err)
	}
}

func TestCT_SharedUserMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_SharedUser()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_SharedUser()
	xml.Unmarshal(buf, v2)
}
