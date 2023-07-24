package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_TablePartsConstructor(t *testing.T) {
	v := sml.NewCT_TableParts()
	if v == nil {
		t.Errorf("sml.NewCT_TableParts must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_TableParts should validate: %s", err)
	}
}

func TestCT_TablePartsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_TableParts()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_TableParts()
	xml.Unmarshal(buf, v2)
}
