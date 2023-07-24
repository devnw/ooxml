package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_DLblsChoiceConstructor(t *testing.T) {
	v := chart.NewCT_DLblsChoice()
	if v == nil {
		t.Errorf("chart.NewCT_DLblsChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_DLblsChoice should validate: %s", err)
	}
}

func TestCT_DLblsChoiceMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_DLblsChoice()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_DLblsChoice()
	xml.Unmarshal(buf, v2)
}
