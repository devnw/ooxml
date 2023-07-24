package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestEG_BarChartSharedConstructor(t *testing.T) {
	v := chart.NewEG_BarChartShared()
	if v == nil {
		t.Errorf("chart.NewEG_BarChartShared must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.EG_BarChartShared should validate: %s", err)
	}
}

func TestEG_BarChartSharedMarshalUnmarshal(t *testing.T) {
	v := chart.NewEG_BarChartShared()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewEG_BarChartShared()
	xml.Unmarshal(buf, v2)
}
