package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestEG_LineChartSharedConstructor(t *testing.T) {
	v := chart.NewEG_LineChartShared()
	if v == nil {
		t.Errorf("chart.NewEG_LineChartShared must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.EG_LineChartShared should validate: %s", err)
	}
}

func TestEG_LineChartSharedMarshalUnmarshal(t *testing.T) {
	v := chart.NewEG_LineChartShared()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewEG_LineChartShared()
	xml.Unmarshal(buf, v2)
}
