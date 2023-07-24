package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_TickMarkConstructor(t *testing.T) {
	v := chart.NewCT_TickMark()
	if v == nil {
		t.Errorf("chart.NewCT_TickMark must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_TickMark should validate: %s", err)
	}
}

func TestCT_TickMarkMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_TickMark()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_TickMark()
	xml.Unmarshal(buf, v2)
}
