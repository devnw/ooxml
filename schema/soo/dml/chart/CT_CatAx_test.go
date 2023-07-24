package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_CatAxConstructor(t *testing.T) {
	v := chart.NewCT_CatAx()
	if v == nil {
		t.Errorf("chart.NewCT_CatAx must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_CatAx should validate: %s", err)
	}
}

func TestCT_CatAxMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_CatAx()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_CatAx()
	xml.Unmarshal(buf, v2)
}
