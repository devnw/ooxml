package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_GapAmountConstructor(t *testing.T) {
	v := chart.NewCT_GapAmount()
	if v == nil {
		t.Errorf("chart.NewCT_GapAmount must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_GapAmount should validate: %s", err)
	}
}

func TestCT_GapAmountMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_GapAmount()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_GapAmount()
	xml.Unmarshal(buf, v2)
}
