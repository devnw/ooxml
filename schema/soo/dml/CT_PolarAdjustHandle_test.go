package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_PolarAdjustHandleConstructor(t *testing.T) {
	v := dml.NewCT_PolarAdjustHandle()
	if v == nil {
		t.Errorf("dml.NewCT_PolarAdjustHandle must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_PolarAdjustHandle should validate: %s", err)
	}
}

func TestCT_PolarAdjustHandleMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_PolarAdjustHandle()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_PolarAdjustHandle()
	xml.Unmarshal(buf, v2)
}
