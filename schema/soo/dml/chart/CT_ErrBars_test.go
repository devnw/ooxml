package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_ErrBarsConstructor(t *testing.T) {
	v := chart.NewCT_ErrBars()
	if v == nil {
		t.Errorf("chart.NewCT_ErrBars must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_ErrBars should validate: %s", err)
	}
}

func TestCT_ErrBarsMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_ErrBars()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_ErrBars()
	xml.Unmarshal(buf, v2)
}
