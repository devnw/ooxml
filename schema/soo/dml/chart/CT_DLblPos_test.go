package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_DLblPosConstructor(t *testing.T) {
	v := chart.NewCT_DLblPos()
	if v == nil {
		t.Errorf("chart.NewCT_DLblPos must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_DLblPos should validate: %s", err)
	}
}

func TestCT_DLblPosMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_DLblPos()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_DLblPos()
	xml.Unmarshal(buf, v2)
}
