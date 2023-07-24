package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_TableStyleInfoConstructor(t *testing.T) {
	v := sml.NewCT_TableStyleInfo()
	if v == nil {
		t.Errorf("sml.NewCT_TableStyleInfo must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_TableStyleInfo should validate: %s", err)
	}
}

func TestCT_TableStyleInfoMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_TableStyleInfo()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_TableStyleInfo()
	xml.Unmarshal(buf, v2)
}
