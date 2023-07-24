package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_DispUnitsChoiceConstructor(t *testing.T) {
	v := chart.NewCT_DispUnitsChoice()
	if v == nil {
		t.Errorf("chart.NewCT_DispUnitsChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_DispUnitsChoice should validate: %s", err)
	}
}

func TestCT_DispUnitsChoiceMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_DispUnitsChoice()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_DispUnitsChoice()
	xml.Unmarshal(buf, v2)
}
