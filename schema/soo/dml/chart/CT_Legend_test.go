package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_LegendConstructor(t *testing.T) {
	v := chart.NewCT_Legend()
	if v == nil {
		t.Errorf("chart.NewCT_Legend must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_Legend should validate: %s", err)
	}
}

func TestCT_LegendMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_Legend()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_Legend()
	xml.Unmarshal(buf, v2)
}
