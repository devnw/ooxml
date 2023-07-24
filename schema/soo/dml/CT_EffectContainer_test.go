package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_EffectContainerConstructor(t *testing.T) {
	v := dml.NewCT_EffectContainer()
	if v == nil {
		t.Errorf("dml.NewCT_EffectContainer must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_EffectContainer should validate: %s", err)
	}
}

func TestCT_EffectContainerMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_EffectContainer()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_EffectContainer()
	xml.Unmarshal(buf, v2)
}
