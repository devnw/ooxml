package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_BarGroupingConstructor(t *testing.T) {
	v := chart.NewCT_BarGrouping()
	if v == nil {
		t.Errorf("chart.NewCT_BarGrouping must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_BarGrouping should validate: %s", err)
	}
}

func TestCT_BarGroupingMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_BarGrouping()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_BarGrouping()
	xml.Unmarshal(buf, v2)
}
