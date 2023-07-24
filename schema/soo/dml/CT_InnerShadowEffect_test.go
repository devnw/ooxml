package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_InnerShadowEffectConstructor(t *testing.T) {
	v := dml.NewCT_InnerShadowEffect()
	if v == nil {
		t.Errorf("dml.NewCT_InnerShadowEffect must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_InnerShadowEffect should validate: %s", err)
	}
}

func TestCT_InnerShadowEffectMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_InnerShadowEffect()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_InnerShadowEffect()
	xml.Unmarshal(buf, v2)
}
