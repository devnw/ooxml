package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_CellStyleConstructor(t *testing.T) {
	v := sml.NewCT_CellStyle()
	if v == nil {
		t.Errorf("sml.NewCT_CellStyle must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_CellStyle should validate: %s", err)
	}
}

func TestCT_CellStyleMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_CellStyle()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_CellStyle()
	xml.Unmarshal(buf, v2)
}
