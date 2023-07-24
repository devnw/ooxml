package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_PivotSourceConstructor(t *testing.T) {
	v := chart.NewCT_PivotSource()
	if v == nil {
		t.Errorf("chart.NewCT_PivotSource must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_PivotSource should validate: %s", err)
	}
}

func TestCT_PivotSourceMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_PivotSource()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_PivotSource()
	xml.Unmarshal(buf, v2)
}
