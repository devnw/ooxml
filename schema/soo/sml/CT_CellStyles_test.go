package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_CellStylesConstructor(t *testing.T) {
	v := sml.NewCT_CellStyles()
	if v == nil {
		t.Errorf("sml.NewCT_CellStyles must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_CellStyles should validate: %s", err)
	}
}

func TestCT_CellStylesMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_CellStyles()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_CellStyles()
	xml.Unmarshal(buf, v2)
}
