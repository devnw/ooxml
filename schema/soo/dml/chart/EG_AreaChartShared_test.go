package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestEG_AreaChartSharedConstructor(t *testing.T) {
	v := chart.NewEG_AreaChartShared()
	if v == nil {
		t.Errorf("chart.NewEG_AreaChartShared must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.EG_AreaChartShared should validate: %s", err)
	}
}

func TestEG_AreaChartSharedMarshalUnmarshal(t *testing.T) {
	v := chart.NewEG_AreaChartShared()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewEG_AreaChartShared()
	xml.Unmarshal(buf, v2)
}
