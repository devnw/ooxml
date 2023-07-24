package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_NumDataSourceChoiceConstructor(t *testing.T) {
	v := chart.NewCT_NumDataSourceChoice()
	if v == nil {
		t.Errorf("chart.NewCT_NumDataSourceChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_NumDataSourceChoice should validate: %s", err)
	}
}

func TestCT_NumDataSourceChoiceMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_NumDataSourceChoice()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_NumDataSourceChoice()
	xml.Unmarshal(buf, v2)
}
