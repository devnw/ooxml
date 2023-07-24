package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_FirstSliceAngConstructor(t *testing.T) {
	v := chart.NewCT_FirstSliceAng()
	if v == nil {
		t.Errorf("chart.NewCT_FirstSliceAng must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_FirstSliceAng should validate: %s", err)
	}
}

func TestCT_FirstSliceAngMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_FirstSliceAng()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_FirstSliceAng()
	xml.Unmarshal(buf, v2)
}
