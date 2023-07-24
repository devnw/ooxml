package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_GroupingConstructor(t *testing.T) {
	v := chart.NewCT_Grouping()
	if v == nil {
		t.Errorf("chart.NewCT_Grouping must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_Grouping should validate: %s", err)
	}
}

func TestCT_GroupingMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_Grouping()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_Grouping()
	xml.Unmarshal(buf, v2)
}
