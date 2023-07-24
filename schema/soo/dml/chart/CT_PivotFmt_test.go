package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_PivotFmtConstructor(t *testing.T) {
	v := chart.NewCT_PivotFmt()
	if v == nil {
		t.Errorf("chart.NewCT_PivotFmt must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_PivotFmt should validate: %s", err)
	}
}

func TestCT_PivotFmtMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_PivotFmt()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_PivotFmt()
	xml.Unmarshal(buf, v2)
}
