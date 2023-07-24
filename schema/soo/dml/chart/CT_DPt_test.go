package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_DPtConstructor(t *testing.T) {
	v := chart.NewCT_DPt()
	if v == nil {
		t.Errorf("chart.NewCT_DPt must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_DPt should validate: %s", err)
	}
}

func TestCT_DPtMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_DPt()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_DPt()
	xml.Unmarshal(buf, v2)
}
