package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_AngleConstructor(t *testing.T) {
	v := dml.NewCT_Angle()
	if v == nil {
		t.Errorf("dml.NewCT_Angle must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_Angle should validate: %s", err)
	}
}

func TestCT_AngleMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_Angle()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_Angle()
	xml.Unmarshal(buf, v2)
}
