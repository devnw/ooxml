package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestEG_EffectPropertiesConstructor(t *testing.T) {
	v := dml.NewEG_EffectProperties()
	if v == nil {
		t.Errorf("dml.NewEG_EffectProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.EG_EffectProperties should validate: %s", err)
	}
}

func TestEG_EffectPropertiesMarshalUnmarshal(t *testing.T) {
	v := dml.NewEG_EffectProperties()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewEG_EffectProperties()
	xml.Unmarshal(buf, v2)
}
