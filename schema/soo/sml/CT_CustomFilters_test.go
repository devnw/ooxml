package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_CustomFiltersConstructor(t *testing.T) {
	v := sml.NewCT_CustomFilters()
	if v == nil {
		t.Errorf("sml.NewCT_CustomFilters must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_CustomFilters should validate: %s", err)
	}
}

func TestCT_CustomFiltersMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_CustomFilters()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_CustomFilters()
	xml.Unmarshal(buf, v2)
}
