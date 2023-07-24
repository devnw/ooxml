package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_TrendlineTypeConstructor(t *testing.T) {
	v := chart.NewCT_TrendlineType()
	if v == nil {
		t.Errorf("chart.NewCT_TrendlineType must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_TrendlineType should validate: %s", err)
	}
}

func TestCT_TrendlineTypeMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_TrendlineType()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_TrendlineType()
	xml.Unmarshal(buf, v2)
}
