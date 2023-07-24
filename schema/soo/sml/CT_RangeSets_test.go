package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_RangeSetsConstructor(t *testing.T) {
	v := sml.NewCT_RangeSets()
	if v == nil {
		t.Errorf("sml.NewCT_RangeSets must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_RangeSets should validate: %s", err)
	}
}

func TestCT_RangeSetsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_RangeSets()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_RangeSets()
	xml.Unmarshal(buf, v2)
}
