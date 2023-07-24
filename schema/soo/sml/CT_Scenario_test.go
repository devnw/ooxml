package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ScenarioConstructor(t *testing.T) {
	v := sml.NewCT_Scenario()
	if v == nil {
		t.Errorf("sml.NewCT_Scenario must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Scenario should validate: %s", err)
	}
}

func TestCT_ScenarioMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Scenario()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Scenario()
	xml.Unmarshal(buf, v2)
}
