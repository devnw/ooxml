package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_StrRefConstructor(t *testing.T) {
	v := chart.NewCT_StrRef()
	if v == nil {
		t.Errorf("chart.NewCT_StrRef must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_StrRef should validate: %s", err)
	}
}

func TestCT_StrRefMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_StrRef()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_StrRef()
	xml.Unmarshal(buf, v2)
}
