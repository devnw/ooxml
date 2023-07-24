package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_UpDownBarConstructor(t *testing.T) {
	v := chart.NewCT_UpDownBar()
	if v == nil {
		t.Errorf("chart.NewCT_UpDownBar must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_UpDownBar should validate: %s", err)
	}
}

func TestCT_UpDownBarMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_UpDownBar()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_UpDownBar()
	xml.Unmarshal(buf, v2)
}
