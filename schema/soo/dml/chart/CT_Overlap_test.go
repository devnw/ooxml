package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_OverlapConstructor(t *testing.T) {
	v := chart.NewCT_Overlap()
	if v == nil {
		t.Errorf("chart.NewCT_Overlap must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_Overlap should validate: %s", err)
	}
}

func TestCT_OverlapMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_Overlap()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_Overlap()
	xml.Unmarshal(buf, v2)
}
