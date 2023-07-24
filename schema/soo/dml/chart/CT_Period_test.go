package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_PeriodConstructor(t *testing.T) {
	v := chart.NewCT_Period()
	if v == nil {
		t.Errorf("chart.NewCT_Period must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_Period should validate: %s", err)
	}
}

func TestCT_PeriodMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_Period()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_Period()
	xml.Unmarshal(buf, v2)
}
