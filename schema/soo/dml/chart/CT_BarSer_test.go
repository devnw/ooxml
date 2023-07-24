package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_BarSerConstructor(t *testing.T) {
	v := chart.NewCT_BarSer()
	if v == nil {
		t.Errorf("chart.NewCT_BarSer must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_BarSer should validate: %s", err)
	}
}

func TestCT_BarSerMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_BarSer()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_BarSer()
	xml.Unmarshal(buf, v2)
}
