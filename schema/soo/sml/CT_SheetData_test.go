package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_SheetDataConstructor(t *testing.T) {
	v := sml.NewCT_SheetData()
	if v == nil {
		t.Errorf("sml.NewCT_SheetData must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_SheetData should validate: %s", err)
	}
}

func TestCT_SheetDataMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_SheetData()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_SheetData()
	xml.Unmarshal(buf, v2)
}
