package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_ExtensionConstructor(t *testing.T) {
	v := chart.NewCT_Extension()
	if v == nil {
		t.Errorf("chart.NewCT_Extension must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_Extension should validate: %s", err)
	}
}

func TestCT_ExtensionMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_Extension()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_Extension()
	xml.Unmarshal(buf, v2)
}
