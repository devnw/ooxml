package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_ValAxConstructor(t *testing.T) {
	v := chart.NewCT_ValAx()
	if v == nil {
		t.Errorf("chart.NewCT_ValAx must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_ValAx should validate: %s", err)
	}
}

func TestCT_ValAxMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_ValAx()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_ValAx()
	xml.Unmarshal(buf, v2)
}
