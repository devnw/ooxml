package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_PivotFiltersConstructor(t *testing.T) {
	v := sml.NewCT_PivotFilters()
	if v == nil {
		t.Errorf("sml.NewCT_PivotFilters must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_PivotFilters should validate: %s", err)
	}
}

func TestCT_PivotFiltersMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_PivotFilters()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_PivotFilters()
	xml.Unmarshal(buf, v2)
}
