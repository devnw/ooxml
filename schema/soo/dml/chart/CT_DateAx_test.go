package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_DateAxConstructor(t *testing.T) {
	v := chart.NewCT_DateAx()
	if v == nil {
		t.Errorf("chart.NewCT_DateAx must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_DateAx should validate: %s", err)
	}
}

func TestCT_DateAxMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_DateAx()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_DateAx()
	xml.Unmarshal(buf, v2)
}
