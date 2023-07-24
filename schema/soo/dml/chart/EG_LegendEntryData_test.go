package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestEG_LegendEntryDataConstructor(t *testing.T) {
	v := chart.NewEG_LegendEntryData()
	if v == nil {
		t.Errorf("chart.NewEG_LegendEntryData must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.EG_LegendEntryData should validate: %s", err)
	}
}

func TestEG_LegendEntryDataMarshalUnmarshal(t *testing.T) {
	v := chart.NewEG_LegendEntryData()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewEG_LegendEntryData()
	xml.Unmarshal(buf, v2)
}
