package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_MarkerConstructor(t *testing.T) {
	v := chart.NewCT_Marker()
	if v == nil {
		t.Errorf("chart.NewCT_Marker must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_Marker should validate: %s", err)
	}
}

func TestCT_MarkerMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_Marker()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_Marker()
	xml.Unmarshal(buf, v2)
}
