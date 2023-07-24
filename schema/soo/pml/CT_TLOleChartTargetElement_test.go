package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TLOleChartTargetElementConstructor(t *testing.T) {
	v := pml.NewCT_TLOleChartTargetElement()
	if v == nil {
		t.Errorf("pml.NewCT_TLOleChartTargetElement must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLOleChartTargetElement should validate: %s", err)
	}
}

func TestCT_TLOleChartTargetElementMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLOleChartTargetElement()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLOleChartTargetElement()
	xml.Unmarshal(buf, v2)
}
