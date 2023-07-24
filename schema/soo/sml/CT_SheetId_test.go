package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_SheetIdConstructor(t *testing.T) {
	v := sml.NewCT_SheetId()
	if v == nil {
		t.Errorf("sml.NewCT_SheetId must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_SheetId should validate: %s", err)
	}
}

func TestCT_SheetIdMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_SheetId()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_SheetId()
	xml.Unmarshal(buf, v2)
}
