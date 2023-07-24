package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_PivotFmtsConstructor(t *testing.T) {
	v := chart.NewCT_PivotFmts()
	if v == nil {
		t.Errorf("chart.NewCT_PivotFmts must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_PivotFmts should validate: %s", err)
	}
}

func TestCT_PivotFmtsMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_PivotFmts()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_PivotFmts()
	xml.Unmarshal(buf, v2)
}
