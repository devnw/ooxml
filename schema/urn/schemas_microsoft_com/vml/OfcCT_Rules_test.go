package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcCT_RulesConstructor(t *testing.T) {
	v := vml.NewOfcCT_Rules()
	if v == nil {
		t.Errorf("vml.NewOfcCT_Rules must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcCT_Rules should validate: %s", err)
	}
}

func TestOfcCT_RulesMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcCT_Rules()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcCT_Rules()
	xml.Unmarshal(buf, v2)
}
