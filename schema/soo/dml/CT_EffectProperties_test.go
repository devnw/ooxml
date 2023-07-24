package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_EffectPropertiesConstructor(t *testing.T) {
	v := dml.NewCT_EffectProperties()
	if v == nil {
		t.Errorf("dml.NewCT_EffectProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_EffectProperties should validate: %s", err)
	}
}

func TestCT_EffectPropertiesMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_EffectProperties()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_EffectProperties()
	xml.Unmarshal(buf, v2)
}
