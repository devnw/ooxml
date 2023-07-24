package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_CellSmartTagPrConstructor(t *testing.T) {
	v := sml.NewCT_CellSmartTagPr()
	if v == nil {
		t.Errorf("sml.NewCT_CellSmartTagPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_CellSmartTagPr should validate: %s", err)
	}
}

func TestCT_CellSmartTagPrMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_CellSmartTagPr()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_CellSmartTagPr()
	xml.Unmarshal(buf, v2)
}
