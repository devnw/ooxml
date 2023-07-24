package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_PivotCacheDefinitionConstructor(t *testing.T) {
	v := sml.NewCT_PivotCacheDefinition()
	if v == nil {
		t.Errorf("sml.NewCT_PivotCacheDefinition must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_PivotCacheDefinition should validate: %s", err)
	}
}

func TestCT_PivotCacheDefinitionMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_PivotCacheDefinition()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_PivotCacheDefinition()
	xml.Unmarshal(buf, v2)
}
