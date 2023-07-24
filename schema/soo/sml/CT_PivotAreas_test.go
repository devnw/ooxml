package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_PivotAreasConstructor(t *testing.T) {
	v := sml.NewCT_PivotAreas()
	if v == nil {
		t.Errorf("sml.NewCT_PivotAreas must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_PivotAreas should validate: %s", err)
	}
}

func TestCT_PivotAreasMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_PivotAreas()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_PivotAreas()
	xml.Unmarshal(buf, v2)
}
