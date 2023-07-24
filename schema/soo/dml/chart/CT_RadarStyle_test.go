package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_RadarStyleConstructor(t *testing.T) {
	v := chart.NewCT_RadarStyle()
	if v == nil {
		t.Errorf("chart.NewCT_RadarStyle must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_RadarStyle should validate: %s", err)
	}
}

func TestCT_RadarStyleMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_RadarStyle()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_RadarStyle()
	xml.Unmarshal(buf, v2)
}
