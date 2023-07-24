package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_DLblsConstructor(t *testing.T) {
	v := chart.NewCT_DLbls()
	if v == nil {
		t.Errorf("chart.NewCT_DLbls must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_DLbls should validate: %s", err)
	}
}

func TestCT_DLblsMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_DLbls()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_DLbls()
	xml.Unmarshal(buf, v2)
}
