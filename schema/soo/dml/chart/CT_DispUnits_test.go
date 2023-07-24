package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_DispUnitsConstructor(t *testing.T) {
	v := chart.NewCT_DispUnits()
	if v == nil {
		t.Errorf("chart.NewCT_DispUnits must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_DispUnits should validate: %s", err)
	}
}

func TestCT_DispUnitsMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_DispUnits()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_DispUnits()
	xml.Unmarshal(buf, v2)
}
