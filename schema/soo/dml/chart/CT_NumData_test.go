package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_NumDataConstructor(t *testing.T) {
	v := chart.NewCT_NumData()
	if v == nil {
		t.Errorf("chart.NewCT_NumData must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_NumData should validate: %s", err)
	}
}

func TestCT_NumDataMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_NumData()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_NumData()
	xml.Unmarshal(buf, v2)
}
