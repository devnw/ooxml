package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_PresetShadowEffectConstructor(t *testing.T) {
	v := dml.NewCT_PresetShadowEffect()
	if v == nil {
		t.Errorf("dml.NewCT_PresetShadowEffect must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_PresetShadowEffect should validate: %s", err)
	}
}

func TestCT_PresetShadowEffectMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_PresetShadowEffect()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_PresetShadowEffect()
	xml.Unmarshal(buf, v2)
}
