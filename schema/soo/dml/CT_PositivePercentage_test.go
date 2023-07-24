package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_PositivePercentageConstructor(t *testing.T) {
	v := dml.NewCT_PositivePercentage()
	if v == nil {
		t.Errorf("dml.NewCT_PositivePercentage must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_PositivePercentage should validate: %s", err)
	}
}

func TestCT_PositivePercentageMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_PositivePercentage()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_PositivePercentage()
	xml.Unmarshal(buf, v2)
}
