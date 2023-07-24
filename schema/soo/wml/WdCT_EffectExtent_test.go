package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestWdCT_EffectExtentConstructor(t *testing.T) {
	v := wml.NewWdCT_EffectExtent()
	if v == nil {
		t.Errorf("wml.NewWdCT_EffectExtent must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.WdCT_EffectExtent should validate: %s", err)
	}
}

func TestWdCT_EffectExtentMarshalUnmarshal(t *testing.T) {
	v := wml.NewWdCT_EffectExtent()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewWdCT_EffectExtent()
	xml.Unmarshal(buf, v2)
}
