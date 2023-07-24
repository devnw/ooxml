package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_FixedPercentageConstructor(t *testing.T) {
	v := dml.NewCT_FixedPercentage()
	if v == nil {
		t.Errorf("dml.NewCT_FixedPercentage must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_FixedPercentage should validate: %s", err)
	}
}

func TestCT_FixedPercentageMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_FixedPercentage()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_FixedPercentage()
	xml.Unmarshal(buf, v2)
}
