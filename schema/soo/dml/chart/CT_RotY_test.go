package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_RotYConstructor(t *testing.T) {
	v := chart.NewCT_RotY()
	if v == nil {
		t.Errorf("chart.NewCT_RotY must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_RotY should validate: %s", err)
	}
}

func TestCT_RotYMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_RotY()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_RotY()
	xml.Unmarshal(buf, v2)
}
