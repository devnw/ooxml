package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestCT_FormulasConstructor(t *testing.T) {
	v := vml.NewCT_Formulas()
	if v == nil {
		t.Errorf("vml.NewCT_Formulas must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.CT_Formulas should validate: %s", err)
	}
}

func TestCT_FormulasMarshalUnmarshal(t *testing.T) {
	v := vml.NewCT_Formulas()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewCT_Formulas()
	xml.Unmarshal(buf, v2)
}
