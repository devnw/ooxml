package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TLAnimateEffectBehaviorConstructor(t *testing.T) {
	v := pml.NewCT_TLAnimateEffectBehavior()
	if v == nil {
		t.Errorf("pml.NewCT_TLAnimateEffectBehavior must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLAnimateEffectBehavior should validate: %s", err)
	}
}

func TestCT_TLAnimateEffectBehaviorMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLAnimateEffectBehavior()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLAnimateEffectBehavior()
	xml.Unmarshal(buf, v2)
}
