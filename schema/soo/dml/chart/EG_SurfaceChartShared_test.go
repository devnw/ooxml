package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestEG_SurfaceChartSharedConstructor(t *testing.T) {
	v := chart.NewEG_SurfaceChartShared()
	if v == nil {
		t.Errorf("chart.NewEG_SurfaceChartShared must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.EG_SurfaceChartShared should validate: %s", err)
	}
}

func TestEG_SurfaceChartSharedMarshalUnmarshal(t *testing.T) {
	v := chart.NewEG_SurfaceChartShared()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewEG_SurfaceChartShared()
	xml.Unmarshal(buf, v2)
}
