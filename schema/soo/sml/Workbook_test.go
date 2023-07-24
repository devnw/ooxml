package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestWorkbookConstructor(t *testing.T) {
	v := sml.NewWorkbook()
	if v == nil {
		t.Errorf("sml.NewWorkbook must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.Workbook should validate: %s", err)
	}
}

func TestWorkbookMarshalUnmarshal(t *testing.T) {
	v := sml.NewWorkbook()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewWorkbook()
	xml.Unmarshal(buf, v2)
}
