package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_EffectReferenceConstructor(t *testing.T) {
	v := dml.NewCT_EffectReference()
	if v == nil {
		t.Errorf("dml.NewCT_EffectReference must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_EffectReference should validate: %s", err)
	}
}

func TestCT_EffectReferenceMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_EffectReference()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_EffectReference()
	xml.Unmarshal(buf, v2)
}
