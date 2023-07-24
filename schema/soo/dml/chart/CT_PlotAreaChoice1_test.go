package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_PlotAreaChoice1Constructor(t *testing.T) {
	v := chart.NewCT_PlotAreaChoice1()
	if v == nil {
		t.Errorf("chart.NewCT_PlotAreaChoice1 must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_PlotAreaChoice1 should validate: %s", err)
	}
}

func TestCT_PlotAreaChoice1MarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_PlotAreaChoice1()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_PlotAreaChoice1()
	xml.Unmarshal(buf, v2)
}
