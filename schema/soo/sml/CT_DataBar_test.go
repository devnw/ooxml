package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_DataBarConstructor(t *testing.T) {
	v := sml.NewCT_DataBar()
	if v == nil {
		t.Errorf("sml.NewCT_DataBar must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_DataBar should validate: %s", err)
	}
}

func TestCT_DataBarMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_DataBar()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_DataBar()
	xml.Unmarshal(buf, v2)
}
