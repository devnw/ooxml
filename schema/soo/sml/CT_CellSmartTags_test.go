package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_CellSmartTagsConstructor(t *testing.T) {
	v := sml.NewCT_CellSmartTags()
	if v == nil {
		t.Errorf("sml.NewCT_CellSmartTags must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_CellSmartTags should validate: %s", err)
	}
}

func TestCT_CellSmartTagsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_CellSmartTags()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_CellSmartTags()
	xml.Unmarshal(buf, v2)
}
