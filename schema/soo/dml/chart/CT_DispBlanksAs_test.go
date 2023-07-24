package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_DispBlanksAsConstructor(t *testing.T) {
	v := chart.NewCT_DispBlanksAs()
	if v == nil {
		t.Errorf("chart.NewCT_DispBlanksAs must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_DispBlanksAs should validate: %s", err)
	}
}

func TestCT_DispBlanksAsMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_DispBlanksAs()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_DispBlanksAs()
	xml.Unmarshal(buf, v2)
}
