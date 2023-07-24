package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_DepthPercentConstructor(t *testing.T) {
	v := chart.NewCT_DepthPercent()
	if v == nil {
		t.Errorf("chart.NewCT_DepthPercent must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_DepthPercent should validate: %s", err)
	}
}

func TestCT_DepthPercentMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_DepthPercent()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_DepthPercent()
	xml.Unmarshal(buf, v2)
}
