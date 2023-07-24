package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestPivotTableDefinitionConstructor(t *testing.T) {
	v := sml.NewPivotTableDefinition()
	if v == nil {
		t.Errorf("sml.NewPivotTableDefinition must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.PivotTableDefinition should validate: %s", err)
	}
}

func TestPivotTableDefinitionMarshalUnmarshal(t *testing.T) {
	v := sml.NewPivotTableDefinition()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewPivotTableDefinition()
	xml.Unmarshal(buf, v2)
}
