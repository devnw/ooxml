package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_AreaSerConstructor(t *testing.T) {
	v := chart.NewCT_AreaSer()
	if v == nil {
		t.Errorf("chart.NewCT_AreaSer must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_AreaSer should validate: %s", err)
	}
}

func TestCT_AreaSerMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_AreaSer()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_AreaSer()
	xml.Unmarshal(buf, v2)
}
