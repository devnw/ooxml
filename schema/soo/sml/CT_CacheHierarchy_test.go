package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_CacheHierarchyConstructor(t *testing.T) {
	v := sml.NewCT_CacheHierarchy()
	if v == nil {
		t.Errorf("sml.NewCT_CacheHierarchy must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_CacheHierarchy should validate: %s", err)
	}
}

func TestCT_CacheHierarchyMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_CacheHierarchy()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_CacheHierarchy()
	xml.Unmarshal(buf, v2)
}
