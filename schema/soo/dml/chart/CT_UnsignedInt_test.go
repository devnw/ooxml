package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_UnsignedIntConstructor(t *testing.T) {
	v := chart.NewCT_UnsignedInt()
	if v == nil {
		t.Errorf("chart.NewCT_UnsignedInt must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_UnsignedInt should validate: %s", err)
	}
}

func TestCT_UnsignedIntMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_UnsignedInt()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_UnsignedInt()
	xml.Unmarshal(buf, v2)
}
