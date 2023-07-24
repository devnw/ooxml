package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_AxDataSourceConstructor(t *testing.T) {
	v := chart.NewCT_AxDataSource()
	if v == nil {
		t.Errorf("chart.NewCT_AxDataSource must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_AxDataSource should validate: %s", err)
	}
}

func TestCT_AxDataSourceMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_AxDataSource()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_AxDataSource()
	xml.Unmarshal(buf, v2)
}
