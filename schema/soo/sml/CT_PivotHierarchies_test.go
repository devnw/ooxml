package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_PivotHierarchiesConstructor(t *testing.T) {
	v := sml.NewCT_PivotHierarchies()
	if v == nil {
		t.Errorf("sml.NewCT_PivotHierarchies must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_PivotHierarchies should validate: %s", err)
	}
}

func TestCT_PivotHierarchiesMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_PivotHierarchies()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_PivotHierarchies()
	xml.Unmarshal(buf, v2)
}
