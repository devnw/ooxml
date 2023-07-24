package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_NumFmtConstructor(t *testing.T) {
	v := chart.NewCT_NumFmt()
	if v == nil {
		t.Errorf("chart.NewCT_NumFmt must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_NumFmt should validate: %s", err)
	}
}

func TestCT_NumFmtMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_NumFmt()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_NumFmt()
	xml.Unmarshal(buf, v2)
}
