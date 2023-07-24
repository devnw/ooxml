package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_PageSetupConstructor(t *testing.T) {
	v := chart.NewCT_PageSetup()
	if v == nil {
		t.Errorf("chart.NewCT_PageSetup must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_PageSetup should validate: %s", err)
	}
}

func TestCT_PageSetupMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_PageSetup()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_PageSetup()
	xml.Unmarshal(buf, v2)
}
