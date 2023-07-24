package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_SpacingRuleConstructor(t *testing.T) {
	v := math.NewCT_SpacingRule()
	if v == nil {
		t.Errorf("math.NewCT_SpacingRule must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_SpacingRule should validate: %s", err)
	}
}

func TestCT_SpacingRuleMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_SpacingRule()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_SpacingRule()
	xml.Unmarshal(buf, v2)
}
