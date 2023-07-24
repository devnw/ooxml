package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_DoubleConstructor(t *testing.T) {
	v := chart.NewCT_Double()
	if v == nil {
		t.Errorf("chart.NewCT_Double must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_Double should validate: %s", err)
	}
}

func TestCT_DoubleMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_Double()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_Double()
	xml.Unmarshal(buf, v2)
}
