package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_GlowEffectConstructor(t *testing.T) {
	v := dml.NewCT_GlowEffect()
	if v == nil {
		t.Errorf("dml.NewCT_GlowEffect must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_GlowEffect should validate: %s", err)
	}
}

func TestCT_GlowEffectMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_GlowEffect()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_GlowEffect()
	xml.Unmarshal(buf, v2)
}
