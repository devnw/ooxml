package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_NumRefConstructor(t *testing.T) {
	v := chart.NewCT_NumRef()
	if v == nil {
		t.Errorf("chart.NewCT_NumRef must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_NumRef should validate: %s", err)
	}
}

func TestCT_NumRefMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_NumRef()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_NumRef()
	xml.Unmarshal(buf, v2)
}
