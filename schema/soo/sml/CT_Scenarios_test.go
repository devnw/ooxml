package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ScenariosConstructor(t *testing.T) {
	v := sml.NewCT_Scenarios()
	if v == nil {
		t.Errorf("sml.NewCT_Scenarios must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Scenarios should validate: %s", err)
	}
}

func TestCT_ScenariosMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Scenarios()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Scenarios()
	xml.Unmarshal(buf, v2)
}
