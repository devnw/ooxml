package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_ErrBarTypeConstructor(t *testing.T) {
	v := chart.NewCT_ErrBarType()
	if v == nil {
		t.Errorf("chart.NewCT_ErrBarType must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_ErrBarType should validate: %s", err)
	}
}

func TestCT_ErrBarTypeMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_ErrBarType()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_ErrBarType()
	xml.Unmarshal(buf, v2)
}
