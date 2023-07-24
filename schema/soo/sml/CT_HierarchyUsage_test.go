package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_HierarchyUsageConstructor(t *testing.T) {
	v := sml.NewCT_HierarchyUsage()
	if v == nil {
		t.Errorf("sml.NewCT_HierarchyUsage must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_HierarchyUsage should validate: %s", err)
	}
}

func TestCT_HierarchyUsageMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_HierarchyUsage()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_HierarchyUsage()
	xml.Unmarshal(buf, v2)
}
