package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_SecondPieSizeConstructor(t *testing.T) {
	v := chart.NewCT_SecondPieSize()
	if v == nil {
		t.Errorf("chart.NewCT_SecondPieSize must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_SecondPieSize should validate: %s", err)
	}
}

func TestCT_SecondPieSizeMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_SecondPieSize()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_SecondPieSize()
	xml.Unmarshal(buf, v2)
}
