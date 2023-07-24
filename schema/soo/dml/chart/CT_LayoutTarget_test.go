package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_LayoutTargetConstructor(t *testing.T) {
	v := chart.NewCT_LayoutTarget()
	if v == nil {
		t.Errorf("chart.NewCT_LayoutTarget must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_LayoutTarget should validate: %s", err)
	}
}

func TestCT_LayoutTargetMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_LayoutTarget()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_LayoutTarget()
	xml.Unmarshal(buf, v2)
}
