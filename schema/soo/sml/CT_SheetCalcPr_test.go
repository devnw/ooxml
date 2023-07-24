package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_SheetCalcPrConstructor(t *testing.T) {
	v := sml.NewCT_SheetCalcPr()
	if v == nil {
		t.Errorf("sml.NewCT_SheetCalcPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_SheetCalcPr should validate: %s", err)
	}
}

func TestCT_SheetCalcPrMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_SheetCalcPr()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_SheetCalcPr()
	xml.Unmarshal(buf, v2)
}
