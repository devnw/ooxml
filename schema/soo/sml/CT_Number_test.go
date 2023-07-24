package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_NumberConstructor(t *testing.T) {
	v := sml.NewCT_Number()
	if v == nil {
		t.Errorf("sml.NewCT_Number must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Number should validate: %s", err)
	}
}

func TestCT_NumberMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Number()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Number()
	xml.Unmarshal(buf, v2)
}
