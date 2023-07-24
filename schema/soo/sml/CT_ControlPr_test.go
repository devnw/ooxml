package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ControlPrConstructor(t *testing.T) {
	v := sml.NewCT_ControlPr()
	if v == nil {
		t.Errorf("sml.NewCT_ControlPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_ControlPr should validate: %s", err)
	}
}

func TestCT_ControlPrMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_ControlPr()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_ControlPr()
	xml.Unmarshal(buf, v2)
}
