package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_PivotHierarchyConstructor(t *testing.T) {
	v := sml.NewCT_PivotHierarchy()
	if v == nil {
		t.Errorf("sml.NewCT_PivotHierarchy must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_PivotHierarchy should validate: %s", err)
	}
}

func TestCT_PivotHierarchyMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_PivotHierarchy()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_PivotHierarchy()
	xml.Unmarshal(buf, v2)
}
