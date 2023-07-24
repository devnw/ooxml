package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_AxDataSourceChoiceConstructor(t *testing.T) {
	v := chart.NewCT_AxDataSourceChoice()
	if v == nil {
		t.Errorf("chart.NewCT_AxDataSourceChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_AxDataSourceChoice should validate: %s", err)
	}
}

func TestCT_AxDataSourceChoiceMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_AxDataSourceChoice()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_AxDataSourceChoice()
	xml.Unmarshal(buf, v2)
}
