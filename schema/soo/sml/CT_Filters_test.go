package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_FiltersConstructor(t *testing.T) {
	v := sml.NewCT_Filters()
	if v == nil {
		t.Errorf("sml.NewCT_Filters must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Filters should validate: %s", err)
	}
}

func TestCT_FiltersMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Filters()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Filters()
	xml.Unmarshal(buf, v2)
}
