package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_SurfaceChartConstructor(t *testing.T) {
	v := chart.NewCT_SurfaceChart()
	if v == nil {
		t.Errorf("chart.NewCT_SurfaceChart must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_SurfaceChart should validate: %s", err)
	}
}

func TestCT_SurfaceChartMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_SurfaceChart()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_SurfaceChart()
	xml.Unmarshal(buf, v2)
}
