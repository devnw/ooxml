package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_FillConstructor(t *testing.T) {
	v := sml.NewCT_Fill()
	if v == nil {
		t.Errorf("sml.NewCT_Fill must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Fill should validate: %s", err)
	}
}

func TestCT_FillMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Fill()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Fill()
	xml.Unmarshal(buf, v2)
}
