package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ProtectedRangesConstructor(t *testing.T) {
	v := sml.NewCT_ProtectedRanges()
	if v == nil {
		t.Errorf("sml.NewCT_ProtectedRanges must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_ProtectedRanges should validate: %s", err)
	}
}

func TestCT_ProtectedRangesMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_ProtectedRanges()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_ProtectedRanges()
	xml.Unmarshal(buf, v2)
}
