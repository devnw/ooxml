package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ConsolidationConstructor(t *testing.T) {
	v := sml.NewCT_Consolidation()
	if v == nil {
		t.Errorf("sml.NewCT_Consolidation must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Consolidation should validate: %s", err)
	}
}

func TestCT_ConsolidationMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Consolidation()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Consolidation()
	xml.Unmarshal(buf, v2)
}
