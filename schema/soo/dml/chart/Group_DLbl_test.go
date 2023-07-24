package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestGroup_DLblConstructor(t *testing.T) {
	v := chart.NewGroup_DLbl()
	if v == nil {
		t.Errorf("chart.NewGroup_DLbl must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.Group_DLbl should validate: %s", err)
	}
}

func TestGroup_DLblMarshalUnmarshal(t *testing.T) {
	v := chart.NewGroup_DLbl()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewGroup_DLbl()
	xml.Unmarshal(buf, v2)
}
