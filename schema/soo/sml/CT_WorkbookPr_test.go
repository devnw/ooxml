package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_WorkbookPrConstructor(t *testing.T) {
	v := sml.NewCT_WorkbookPr()
	if v == nil {
		t.Errorf("sml.NewCT_WorkbookPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_WorkbookPr should validate: %s", err)
	}
}

func TestCT_WorkbookPrMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_WorkbookPr()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_WorkbookPr()
	xml.Unmarshal(buf, v2)
}
