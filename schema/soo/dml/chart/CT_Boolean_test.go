package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_BooleanConstructor(t *testing.T) {
	v := chart.NewCT_Boolean()
	if v == nil {
		t.Errorf("chart.NewCT_Boolean must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_Boolean should validate: %s", err)
	}
}

func TestCT_BooleanMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_Boolean()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_Boolean()
	xml.Unmarshal(buf, v2)
}
