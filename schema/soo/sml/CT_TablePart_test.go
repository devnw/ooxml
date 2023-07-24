package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_TablePartConstructor(t *testing.T) {
	v := sml.NewCT_TablePart()
	if v == nil {
		t.Errorf("sml.NewCT_TablePart must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_TablePart should validate: %s", err)
	}
}

func TestCT_TablePartMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_TablePart()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_TablePart()
	xml.Unmarshal(buf, v2)
}
