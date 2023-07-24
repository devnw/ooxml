package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_ScalingConstructor(t *testing.T) {
	v := chart.NewCT_Scaling()
	if v == nil {
		t.Errorf("chart.NewCT_Scaling must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_Scaling should validate: %s", err)
	}
}

func TestCT_ScalingMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_Scaling()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_Scaling()
	xml.Unmarshal(buf, v2)
}
