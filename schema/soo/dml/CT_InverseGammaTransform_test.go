package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_InverseGammaTransformConstructor(t *testing.T) {
	v := dml.NewCT_InverseGammaTransform()
	if v == nil {
		t.Errorf("dml.NewCT_InverseGammaTransform must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_InverseGammaTransform should validate: %s", err)
	}
}

func TestCT_InverseGammaTransformMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_InverseGammaTransform()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_InverseGammaTransform()
	xml.Unmarshal(buf, v2)
}
