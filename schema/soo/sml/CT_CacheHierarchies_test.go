package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_CacheHierarchiesConstructor(t *testing.T) {
	v := sml.NewCT_CacheHierarchies()
	if v == nil {
		t.Errorf("sml.NewCT_CacheHierarchies must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_CacheHierarchies should validate: %s", err)
	}
}

func TestCT_CacheHierarchiesMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_CacheHierarchies()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_CacheHierarchies()
	xml.Unmarshal(buf, v2)
}
