package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_AxPosConstructor(t *testing.T) {
	v := chart.NewCT_AxPos()
	if v == nil {
		t.Errorf("chart.NewCT_AxPos must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_AxPos should validate: %s", err)
	}
}

func TestCT_AxPosMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_AxPos()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_AxPos()
	xml.Unmarshal(buf, v2)
}
