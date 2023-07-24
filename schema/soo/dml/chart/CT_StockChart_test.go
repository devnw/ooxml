package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_StockChartConstructor(t *testing.T) {
	v := chart.NewCT_StockChart()
	if v == nil {
		t.Errorf("chart.NewCT_StockChart must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_StockChart should validate: %s", err)
	}
}

func TestCT_StockChartMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_StockChart()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_StockChart()
	xml.Unmarshal(buf, v2)
}
