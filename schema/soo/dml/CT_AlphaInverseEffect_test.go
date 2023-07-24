package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_AlphaInverseEffectConstructor(t *testing.T) {
	v := dml.NewCT_AlphaInverseEffect()
	if v == nil {
		t.Errorf("dml.NewCT_AlphaInverseEffect must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_AlphaInverseEffect should validate: %s", err)
	}
}

func TestCT_AlphaInverseEffectMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_AlphaInverseEffect()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_AlphaInverseEffect()
	xml.Unmarshal(buf, v2)
}
