package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_CustSplitConstructor(t *testing.T) {
	v := chart.NewCT_CustSplit()
	if v == nil {
		t.Errorf("chart.NewCT_CustSplit must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_CustSplit should validate: %s", err)
	}
}

func TestCT_CustSplitMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_CustSplit()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_CustSplit()
	xml.Unmarshal(buf, v2)
}
