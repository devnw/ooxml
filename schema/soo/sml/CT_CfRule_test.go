package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_CfRuleConstructor(t *testing.T) {
	v := sml.NewCT_CfRule()
	if v == nil {
		t.Errorf("sml.NewCT_CfRule must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_CfRule should validate: %s", err)
	}
}

func TestCT_CfRuleMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_CfRule()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_CfRule()
	xml.Unmarshal(buf, v2)
}
