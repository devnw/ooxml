package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_RevisionMoveConstructor(t *testing.T) {
	v := sml.NewCT_RevisionMove()
	if v == nil {
		t.Errorf("sml.NewCT_RevisionMove must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_RevisionMove should validate: %s", err)
	}
}

func TestCT_RevisionMoveMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_RevisionMove()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_RevisionMove()
	xml.Unmarshal(buf, v2)
}
