package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_SheetFormatPrConstructor(t *testing.T) {
	v := sml.NewCT_SheetFormatPr()
	if v == nil {
		t.Errorf("sml.NewCT_SheetFormatPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_SheetFormatPr should validate: %s", err)
	}
}

func TestCT_SheetFormatPrMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_SheetFormatPr()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_SheetFormatPr()
	xml.Unmarshal(buf, v2)
}
