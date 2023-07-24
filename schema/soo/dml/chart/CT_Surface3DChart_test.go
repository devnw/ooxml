package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_Surface3DChartConstructor(t *testing.T) {
	v := chart.NewCT_Surface3DChart()
	if v == nil {
		t.Errorf("chart.NewCT_Surface3DChart must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_Surface3DChart should validate: %s", err)
	}
}

func TestCT_Surface3DChartMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_Surface3DChart()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_Surface3DChart()
	xml.Unmarshal(buf, v2)
}
