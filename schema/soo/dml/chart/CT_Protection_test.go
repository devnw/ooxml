package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_ProtectionConstructor(t *testing.T) {
	v := chart.NewCT_Protection()
	if v == nil {
		t.Errorf("chart.NewCT_Protection must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_Protection should validate: %s", err)
	}
}

func TestCT_ProtectionMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_Protection()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_Protection()
	xml.Unmarshal(buf, v2)
}
