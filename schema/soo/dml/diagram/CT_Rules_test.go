package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_RulesConstructor(t *testing.T) {
	v := diagram.NewCT_Rules()
	if v == nil {
		t.Errorf("diagram.NewCT_Rules must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_Rules should validate: %s", err)
	}
}

func TestCT_RulesMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_Rules()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_Rules()
	xml.Unmarshal(buf, v2)
}
