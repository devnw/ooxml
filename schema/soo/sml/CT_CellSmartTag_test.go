package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_CellSmartTagConstructor(t *testing.T) {
	v := sml.NewCT_CellSmartTag()
	if v == nil {
		t.Errorf("sml.NewCT_CellSmartTag must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_CellSmartTag should validate: %s", err)
	}
}

func TestCT_CellSmartTagMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_CellSmartTag()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_CellSmartTag()
	xml.Unmarshal(buf, v2)
}
