package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestPivotCacheDefinitionConstructor(t *testing.T) {
	v := sml.NewPivotCacheDefinition()
	if v == nil {
		t.Errorf("sml.NewPivotCacheDefinition must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.PivotCacheDefinition should validate: %s", err)
	}
}

func TestPivotCacheDefinitionMarshalUnmarshal(t *testing.T) {
	v := sml.NewPivotCacheDefinition()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewPivotCacheDefinition()
	xml.Unmarshal(buf, v2)
}
