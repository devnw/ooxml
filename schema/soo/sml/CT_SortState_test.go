package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_SortStateConstructor(t *testing.T) {
	v := sml.NewCT_SortState()
	if v == nil {
		t.Errorf("sml.NewCT_SortState must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_SortState should validate: %s", err)
	}
}

func TestCT_SortStateMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_SortState()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_SortState()
	xml.Unmarshal(buf, v2)
}
