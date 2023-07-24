package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_BreakConstructor(t *testing.T) {
	v := sml.NewCT_Break()
	if v == nil {
		t.Errorf("sml.NewCT_Break must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Break should validate: %s", err)
	}
}

func TestCT_BreakMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Break()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Break()
	xml.Unmarshal(buf, v2)
}
