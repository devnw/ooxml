package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_SheetPrConstructor(t *testing.T) {
	v := sml.NewCT_SheetPr()
	if v == nil {
		t.Errorf("sml.NewCT_SheetPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_SheetPr should validate: %s", err)
	}
}

func TestCT_SheetPrMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_SheetPr()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_SheetPr()
	xml.Unmarshal(buf, v2)
}
