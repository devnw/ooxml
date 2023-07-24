package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_SplitTypeConstructor(t *testing.T) {
	v := chart.NewCT_SplitType()
	if v == nil {
		t.Errorf("chart.NewCT_SplitType must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_SplitType should validate: %s", err)
	}
}

func TestCT_SplitTypeMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_SplitType()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_SplitType()
	xml.Unmarshal(buf, v2)
}
