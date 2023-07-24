package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_TitleConstructor(t *testing.T) {
	v := chart.NewCT_Title()
	if v == nil {
		t.Errorf("chart.NewCT_Title must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_Title should validate: %s", err)
	}
}

func TestCT_TitleMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_Title()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_Title()
	xml.Unmarshal(buf, v2)
}
