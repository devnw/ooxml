package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_ErrValTypeConstructor(t *testing.T) {
	v := chart.NewCT_ErrValType()
	if v == nil {
		t.Errorf("chart.NewCT_ErrValType must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_ErrValType should validate: %s", err)
	}
}

func TestCT_ErrValTypeMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_ErrValType()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_ErrValType()
	xml.Unmarshal(buf, v2)
}
