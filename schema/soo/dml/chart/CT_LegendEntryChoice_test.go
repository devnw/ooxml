package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_LegendEntryChoiceConstructor(t *testing.T) {
	v := chart.NewCT_LegendEntryChoice()
	if v == nil {
		t.Errorf("chart.NewCT_LegendEntryChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_LegendEntryChoice should validate: %s", err)
	}
}

func TestCT_LegendEntryChoiceMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_LegendEntryChoice()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_LegendEntryChoice()
	xml.Unmarshal(buf, v2)
}
