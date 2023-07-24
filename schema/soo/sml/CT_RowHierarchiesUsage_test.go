package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_RowHierarchiesUsageConstructor(t *testing.T) {
	v := sml.NewCT_RowHierarchiesUsage()
	if v == nil {
		t.Errorf("sml.NewCT_RowHierarchiesUsage must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_RowHierarchiesUsage should validate: %s", err)
	}
}

func TestCT_RowHierarchiesUsageMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_RowHierarchiesUsage()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_RowHierarchiesUsage()
	xml.Unmarshal(buf, v2)
}
