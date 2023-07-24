package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_GammaTransformConstructor(t *testing.T) {
	v := dml.NewCT_GammaTransform()
	if v == nil {
		t.Errorf("dml.NewCT_GammaTransform must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_GammaTransform should validate: %s", err)
	}
}

func TestCT_GammaTransformMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_GammaTransform()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_GammaTransform()
	xml.Unmarshal(buf, v2)
}
