package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_OfPieTypeConstructor(t *testing.T) {
	v := chart.NewCT_OfPieType()
	if v == nil {
		t.Errorf("chart.NewCT_OfPieType must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_OfPieType should validate: %s", err)
	}
}

func TestCT_OfPieTypeMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_OfPieType()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_OfPieType()
	xml.Unmarshal(buf, v2)
}
