package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_PositiveFixedPercentageConstructor(t *testing.T) {
	v := dml.NewCT_PositiveFixedPercentage()
	if v == nil {
		t.Errorf("dml.NewCT_PositiveFixedPercentage must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_PositiveFixedPercentage should validate: %s", err)
	}
}

func TestCT_PositiveFixedPercentageMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_PositiveFixedPercentage()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_PositiveFixedPercentage()
	xml.Unmarshal(buf, v2)
}
