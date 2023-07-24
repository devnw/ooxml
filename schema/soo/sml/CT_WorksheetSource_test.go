package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_WorksheetSourceConstructor(t *testing.T) {
	v := sml.NewCT_WorksheetSource()
	if v == nil {
		t.Errorf("sml.NewCT_WorksheetSource must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_WorksheetSource should validate: %s", err)
	}
}

func TestCT_WorksheetSourceMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_WorksheetSource()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_WorksheetSource()
	xml.Unmarshal(buf, v2)
}
