package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_AxisUnitConstructor(t *testing.T) {
	v := chart.NewCT_AxisUnit()
	if v == nil {
		t.Errorf("chart.NewCT_AxisUnit must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_AxisUnit should validate: %s", err)
	}
}

func TestCT_AxisUnitMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_AxisUnit()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_AxisUnit()
	xml.Unmarshal(buf, v2)
}
