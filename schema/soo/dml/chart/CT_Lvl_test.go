package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_LvlConstructor(t *testing.T) {
	v := chart.NewCT_Lvl()
	if v == nil {
		t.Errorf("chart.NewCT_Lvl must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_Lvl should validate: %s", err)
	}
}

func TestCT_LvlMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_Lvl()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_Lvl()
	xml.Unmarshal(buf, v2)
}
