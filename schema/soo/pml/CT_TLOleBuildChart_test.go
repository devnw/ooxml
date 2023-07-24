package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TLOleBuildChartConstructor(t *testing.T) {
	v := pml.NewCT_TLOleBuildChart()
	if v == nil {
		t.Errorf("pml.NewCT_TLOleBuildChart must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLOleBuildChart should validate: %s", err)
	}
}

func TestCT_TLOleBuildChartMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLOleBuildChart()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLOleBuildChart()
	xml.Unmarshal(buf, v2)
}
