package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_OuterShadowEffectConstructor(t *testing.T) {
	v := dml.NewCT_OuterShadowEffect()
	if v == nil {
		t.Errorf("dml.NewCT_OuterShadowEffect must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_OuterShadowEffect should validate: %s", err)
	}
}

func TestCT_OuterShadowEffectMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_OuterShadowEffect()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_OuterShadowEffect()
	xml.Unmarshal(buf, v2)
}
