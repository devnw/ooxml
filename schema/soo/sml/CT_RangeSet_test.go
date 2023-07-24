package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_RangeSetConstructor(t *testing.T) {
	v := sml.NewCT_RangeSet()
	if v == nil {
		t.Errorf("sml.NewCT_RangeSet must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_RangeSet should validate: %s", err)
	}
}

func TestCT_RangeSetMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_RangeSet()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_RangeSet()
	xml.Unmarshal(buf, v2)
}
