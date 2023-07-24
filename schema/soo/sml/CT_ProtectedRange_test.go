package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ProtectedRangeConstructor(t *testing.T) {
	v := sml.NewCT_ProtectedRange()
	if v == nil {
		t.Errorf("sml.NewCT_ProtectedRange must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_ProtectedRange should validate: %s", err)
	}
}

func TestCT_ProtectedRangeMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_ProtectedRange()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_ProtectedRange()
	xml.Unmarshal(buf, v2)
}
