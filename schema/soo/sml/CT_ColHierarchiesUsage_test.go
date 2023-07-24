package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ColHierarchiesUsageConstructor(t *testing.T) {
	v := sml.NewCT_ColHierarchiesUsage()
	if v == nil {
		t.Errorf("sml.NewCT_ColHierarchiesUsage must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_ColHierarchiesUsage should validate: %s", err)
	}
}

func TestCT_ColHierarchiesUsageMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_ColHierarchiesUsage()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_ColHierarchiesUsage()
	xml.Unmarshal(buf, v2)
}
