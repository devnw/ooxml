package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_TrendlineConstructor(t *testing.T) {
	v := chart.NewCT_Trendline()
	if v == nil {
		t.Errorf("chart.NewCT_Trendline must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_Trendline should validate: %s", err)
	}
}

func TestCT_TrendlineMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_Trendline()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_Trendline()
	xml.Unmarshal(buf, v2)
}
