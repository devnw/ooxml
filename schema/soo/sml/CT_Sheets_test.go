package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_SheetsConstructor(t *testing.T) {
	v := sml.NewCT_Sheets()
	if v == nil {
		t.Errorf("sml.NewCT_Sheets must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Sheets should validate: %s", err)
	}
}

func TestCT_SheetsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Sheets()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Sheets()
	xml.Unmarshal(buf, v2)
}
