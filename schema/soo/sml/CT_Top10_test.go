package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_Top10Constructor(t *testing.T) {
	v := sml.NewCT_Top10()
	if v == nil {
		t.Errorf("sml.NewCT_Top10 must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Top10 should validate: %s", err)
	}
}

func TestCT_Top10MarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Top10()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Top10()
	xml.Unmarshal(buf, v2)
}
