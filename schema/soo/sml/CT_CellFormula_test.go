package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_CellFormulaConstructor(t *testing.T) {
	v := sml.NewCT_CellFormula()
	if v == nil {
		t.Errorf("sml.NewCT_CellFormula must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_CellFormula should validate: %s", err)
	}
}

func TestCT_CellFormulaMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_CellFormula()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_CellFormula()
	xml.Unmarshal(buf, v2)
}
