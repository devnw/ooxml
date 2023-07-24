package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_UpDownBarsConstructor(t *testing.T) {
	v := chart.NewCT_UpDownBars()
	if v == nil {
		t.Errorf("chart.NewCT_UpDownBars must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_UpDownBars should validate: %s", err)
	}
}

func TestCT_UpDownBarsMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_UpDownBars()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_UpDownBars()
	xml.Unmarshal(buf, v2)
}
