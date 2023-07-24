package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_MeasureGroupConstructor(t *testing.T) {
	v := sml.NewCT_MeasureGroup()
	if v == nil {
		t.Errorf("sml.NewCT_MeasureGroup must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_MeasureGroup should validate: %s", err)
	}
}

func TestCT_MeasureGroupMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_MeasureGroup()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_MeasureGroup()
	xml.Unmarshal(buf, v2)
}
