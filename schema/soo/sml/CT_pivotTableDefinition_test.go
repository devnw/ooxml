package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_pivotTableDefinitionConstructor(t *testing.T) {
	v := sml.NewCT_pivotTableDefinition()
	if v == nil {
		t.Errorf("sml.NewCT_pivotTableDefinition must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_pivotTableDefinition should validate: %s", err)
	}
}

func TestCT_pivotTableDefinitionMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_pivotTableDefinition()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_pivotTableDefinition()
	xml.Unmarshal(buf, v2)
}
