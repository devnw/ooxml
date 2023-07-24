package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_WorkbookConstructor(t *testing.T) {
	v := sml.NewCT_Workbook()
	if v == nil {
		t.Errorf("sml.NewCT_Workbook must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Workbook should validate: %s", err)
	}
}

func TestCT_WorkbookMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Workbook()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Workbook()
	xml.Unmarshal(buf, v2)
}
