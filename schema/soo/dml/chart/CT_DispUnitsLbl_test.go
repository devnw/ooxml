package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_DispUnitsLblConstructor(t *testing.T) {
	v := chart.NewCT_DispUnitsLbl()
	if v == nil {
		t.Errorf("chart.NewCT_DispUnitsLbl must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_DispUnitsLbl should validate: %s", err)
	}
}

func TestCT_DispUnitsLblMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_DispUnitsLbl()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_DispUnitsLbl()
	xml.Unmarshal(buf, v2)
}
