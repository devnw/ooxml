package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_FieldGroupConstructor(t *testing.T) {
	v := sml.NewCT_FieldGroup()
	if v == nil {
		t.Errorf("sml.NewCT_FieldGroup must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_FieldGroup should validate: %s", err)
	}
}

func TestCT_FieldGroupMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_FieldGroup()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_FieldGroup()
	xml.Unmarshal(buf, v2)
}
