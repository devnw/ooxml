package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_BubbleScaleConstructor(t *testing.T) {
	v := chart.NewCT_BubbleScale()
	if v == nil {
		t.Errorf("chart.NewCT_BubbleScale must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_BubbleScale should validate: %s", err)
	}
}

func TestCT_BubbleScaleMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_BubbleScale()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_BubbleScale()
	xml.Unmarshal(buf, v2)
}
