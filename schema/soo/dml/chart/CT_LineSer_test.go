package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_LineSerConstructor(t *testing.T) {
	v := chart.NewCT_LineSer()
	if v == nil {
		t.Errorf("chart.NewCT_LineSer must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_LineSer should validate: %s", err)
	}
}

func TestCT_LineSerMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_LineSer()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_LineSer()
	xml.Unmarshal(buf, v2)
}
