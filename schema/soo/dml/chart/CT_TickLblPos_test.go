package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_TickLblPosConstructor(t *testing.T) {
	v := chart.NewCT_TickLblPos()
	if v == nil {
		t.Errorf("chart.NewCT_TickLblPos must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_TickLblPos should validate: %s", err)
	}
}

func TestCT_TickLblPosMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_TickLblPos()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_TickLblPos()
	xml.Unmarshal(buf, v2)
}
