package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_PlotAreaChoiceConstructor(t *testing.T) {
	v := chart.NewCT_PlotAreaChoice()
	if v == nil {
		t.Errorf("chart.NewCT_PlotAreaChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_PlotAreaChoice should validate: %s", err)
	}
}

func TestCT_PlotAreaChoiceMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_PlotAreaChoice()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_PlotAreaChoice()
	xml.Unmarshal(buf, v2)
}
