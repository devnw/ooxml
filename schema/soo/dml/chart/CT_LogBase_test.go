package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_LogBaseConstructor(t *testing.T) {
	v := chart.NewCT_LogBase()
	if v == nil {
		t.Errorf("chart.NewCT_LogBase must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_LogBase should validate: %s", err)
	}
}

func TestCT_LogBaseMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_LogBase()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_LogBase()
	xml.Unmarshal(buf, v2)
}
