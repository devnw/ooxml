package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_LblAlgnConstructor(t *testing.T) {
	v := chart.NewCT_LblAlgn()
	if v == nil {
		t.Errorf("chart.NewCT_LblAlgn must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_LblAlgn should validate: %s", err)
	}
}

func TestCT_LblAlgnMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_LblAlgn()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_LblAlgn()
	xml.Unmarshal(buf, v2)
}
