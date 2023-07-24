package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_HPercentConstructor(t *testing.T) {
	v := chart.NewCT_HPercent()
	if v == nil {
		t.Errorf("chart.NewCT_HPercent must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_HPercent should validate: %s", err)
	}
}

func TestCT_HPercentMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_HPercent()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_HPercent()
	xml.Unmarshal(buf, v2)
}
