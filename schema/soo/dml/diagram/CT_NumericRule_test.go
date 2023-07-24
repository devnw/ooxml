package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_NumericRuleConstructor(t *testing.T) {
	v := diagram.NewCT_NumericRule()
	if v == nil {
		t.Errorf("diagram.NewCT_NumericRule must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_NumericRule should validate: %s", err)
	}
}

func TestCT_NumericRuleMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_NumericRule()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_NumericRule()
	xml.Unmarshal(buf, v2)
}
