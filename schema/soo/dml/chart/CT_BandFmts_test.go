package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_BandFmtsConstructor(t *testing.T) {
	v := chart.NewCT_BandFmts()
	if v == nil {
		t.Errorf("chart.NewCT_BandFmts must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_BandFmts should validate: %s", err)
	}
}

func TestCT_BandFmtsMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_BandFmts()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_BandFmts()
	xml.Unmarshal(buf, v2)
}
