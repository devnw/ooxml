package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_StrValConstructor(t *testing.T) {
	v := chart.NewCT_StrVal()
	if v == nil {
		t.Errorf("chart.NewCT_StrVal must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_StrVal should validate: %s", err)
	}
}

func TestCT_StrValMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_StrVal()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_StrVal()
	xml.Unmarshal(buf, v2)
}
