package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_TrendlineLblConstructor(t *testing.T) {
	v := chart.NewCT_TrendlineLbl()
	if v == nil {
		t.Errorf("chart.NewCT_TrendlineLbl must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_TrendlineLbl should validate: %s", err)
	}
}

func TestCT_TrendlineLblMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_TrendlineLbl()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_TrendlineLbl()
	xml.Unmarshal(buf, v2)
}
