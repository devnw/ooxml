package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_BuiltInUnitConstructor(t *testing.T) {
	v := chart.NewCT_BuiltInUnit()
	if v == nil {
		t.Errorf("chart.NewCT_BuiltInUnit must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_BuiltInUnit should validate: %s", err)
	}
}

func TestCT_BuiltInUnitMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_BuiltInUnit()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_BuiltInUnit()
	xml.Unmarshal(buf, v2)
}
