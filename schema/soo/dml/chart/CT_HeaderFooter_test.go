package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_HeaderFooterConstructor(t *testing.T) {
	v := chart.NewCT_HeaderFooter()
	if v == nil {
		t.Errorf("chart.NewCT_HeaderFooter must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_HeaderFooter should validate: %s", err)
	}
}

func TestCT_HeaderFooterMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_HeaderFooter()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_HeaderFooter()
	xml.Unmarshal(buf, v2)
}
