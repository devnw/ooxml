package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_ColorChangeEffectConstructor(t *testing.T) {
	v := dml.NewCT_ColorChangeEffect()
	if v == nil {
		t.Errorf("dml.NewCT_ColorChangeEffect must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_ColorChangeEffect should validate: %s", err)
	}
}

func TestCT_ColorChangeEffectMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_ColorChangeEffect()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_ColorChangeEffect()
	xml.Unmarshal(buf, v2)
}
