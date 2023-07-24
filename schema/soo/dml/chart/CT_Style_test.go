package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_StyleConstructor(t *testing.T) {
	v := chart.NewCT_Style()
	if v == nil {
		t.Errorf("chart.NewCT_Style must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_Style should validate: %s", err)
	}
}

func TestCT_StyleMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_Style()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_Style()
	xml.Unmarshal(buf, v2)
}
