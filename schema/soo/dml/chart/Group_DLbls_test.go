package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestGroup_DLblsConstructor(t *testing.T) {
	v := chart.NewGroup_DLbls()
	if v == nil {
		t.Errorf("chart.NewGroup_DLbls must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.Group_DLbls should validate: %s", err)
	}
}

func TestGroup_DLblsMarshalUnmarshal(t *testing.T) {
	v := chart.NewGroup_DLbls()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewGroup_DLbls()
	xml.Unmarshal(buf, v2)
}
