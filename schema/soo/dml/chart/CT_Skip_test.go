package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_SkipConstructor(t *testing.T) {
	v := chart.NewCT_Skip()
	if v == nil {
		t.Errorf("chart.NewCT_Skip must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_Skip should validate: %s", err)
	}
}

func TestCT_SkipMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_Skip()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_Skip()
	xml.Unmarshal(buf, v2)
}
