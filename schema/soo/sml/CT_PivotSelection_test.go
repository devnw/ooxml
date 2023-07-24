package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_PivotSelectionConstructor(t *testing.T) {
	v := sml.NewCT_PivotSelection()
	if v == nil {
		t.Errorf("sml.NewCT_PivotSelection must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_PivotSelection should validate: %s", err)
	}
}

func TestCT_PivotSelectionMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_PivotSelection()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_PivotSelection()
	xml.Unmarshal(buf, v2)
}
