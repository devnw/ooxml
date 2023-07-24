package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestFormulasConstructor(t *testing.T) {
	v := vml.NewFormulas()
	if v == nil {
		t.Errorf("vml.NewFormulas must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.Formulas should validate: %s", err)
	}
}

func TestFormulasMarshalUnmarshal(t *testing.T) {
	v := vml.NewFormulas()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewFormulas()
	xml.Unmarshal(buf, v2)
}
