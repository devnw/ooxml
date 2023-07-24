package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_Path2DCubicBezierToConstructor(t *testing.T) {
	v := dml.NewCT_Path2DCubicBezierTo()
	if v == nil {
		t.Errorf("dml.NewCT_Path2DCubicBezierTo must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_Path2DCubicBezierTo should validate: %s", err)
	}
}

func TestCT_Path2DCubicBezierToMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_Path2DCubicBezierTo()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_Path2DCubicBezierTo()
	xml.Unmarshal(buf, v2)
}
