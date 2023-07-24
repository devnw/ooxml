package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_AutoSortScopeConstructor(t *testing.T) {
	v := sml.NewCT_AutoSortScope()
	if v == nil {
		t.Errorf("sml.NewCT_AutoSortScope must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_AutoSortScope should validate: %s", err)
	}
}

func TestCT_AutoSortScopeMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_AutoSortScope()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_AutoSortScope()
	xml.Unmarshal(buf, v2)
}
