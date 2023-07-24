package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TLIterateIntervalPercentageConstructor(t *testing.T) {
	v := pml.NewCT_TLIterateIntervalPercentage()
	if v == nil {
		t.Errorf("pml.NewCT_TLIterateIntervalPercentage must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLIterateIntervalPercentage should validate: %s", err)
	}
}

func TestCT_TLIterateIntervalPercentageMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLIterateIntervalPercentage()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLIterateIntervalPercentage()
	xml.Unmarshal(buf, v2)
}
