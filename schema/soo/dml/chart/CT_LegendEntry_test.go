package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_LegendEntryConstructor(t *testing.T) {
	v := chart.NewCT_LegendEntry()
	if v == nil {
		t.Errorf("chart.NewCT_LegendEntry must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_LegendEntry should validate: %s", err)
	}
}

func TestCT_LegendEntryMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_LegendEntry()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_LegendEntry()
	xml.Unmarshal(buf, v2)
}
