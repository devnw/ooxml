package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_OrderConstructor(t *testing.T) {
	v := chart.NewCT_Order()
	if v == nil {
		t.Errorf("chart.NewCT_Order must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_Order should validate: %s", err)
	}
}

func TestCT_OrderMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_Order()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_Order()
	xml.Unmarshal(buf, v2)
}
