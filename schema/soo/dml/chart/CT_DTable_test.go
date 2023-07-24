package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_DTableConstructor(t *testing.T) {
	v := chart.NewCT_DTable()
	if v == nil {
		t.Errorf("chart.NewCT_DTable must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_DTable should validate: %s", err)
	}
}

func TestCT_DTableMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_DTable()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_DTable()
	xml.Unmarshal(buf, v2)
}
