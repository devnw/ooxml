package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_ExternalDataConstructor(t *testing.T) {
	v := chart.NewCT_ExternalData()
	if v == nil {
		t.Errorf("chart.NewCT_ExternalData must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_ExternalData should validate: %s", err)
	}
}

func TestCT_ExternalDataMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_ExternalData()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_ExternalData()
	xml.Unmarshal(buf, v2)
}
