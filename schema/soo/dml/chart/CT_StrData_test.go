package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_StrDataConstructor(t *testing.T) {
	v := chart.NewCT_StrData()
	if v == nil {
		t.Errorf("chart.NewCT_StrData must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_StrData should validate: %s", err)
	}
}

func TestCT_StrDataMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_StrData()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_StrData()
	xml.Unmarshal(buf, v2)
}
