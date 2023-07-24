package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_TableFormulaConstructor(t *testing.T) {
	v := sml.NewCT_TableFormula()
	if v == nil {
		t.Errorf("sml.NewCT_TableFormula must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_TableFormula should validate: %s", err)
	}
}

func TestCT_TableFormulaMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_TableFormula()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_TableFormula()
	xml.Unmarshal(buf, v2)
}
