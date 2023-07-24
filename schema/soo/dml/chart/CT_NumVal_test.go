package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_NumValConstructor(t *testing.T) {
	v := chart.NewCT_NumVal()
	if v == nil {
		t.Errorf("chart.NewCT_NumVal must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_NumVal should validate: %s", err)
	}
}

func TestCT_NumValMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_NumVal()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_NumVal()
	xml.Unmarshal(buf, v2)
}
