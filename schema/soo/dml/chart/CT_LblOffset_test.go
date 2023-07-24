package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_LblOffsetConstructor(t *testing.T) {
	v := chart.NewCT_LblOffset()
	if v == nil {
		t.Errorf("chart.NewCT_LblOffset must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_LblOffset should validate: %s", err)
	}
}

func TestCT_LblOffsetMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_LblOffset()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_LblOffset()
	xml.Unmarshal(buf, v2)
}
