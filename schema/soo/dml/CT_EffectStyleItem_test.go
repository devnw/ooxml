package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_EffectStyleItemConstructor(t *testing.T) {
	v := dml.NewCT_EffectStyleItem()
	if v == nil {
		t.Errorf("dml.NewCT_EffectStyleItem must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_EffectStyleItem should validate: %s", err)
	}
}

func TestCT_EffectStyleItemMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_EffectStyleItem()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_EffectStyleItem()
	xml.Unmarshal(buf, v2)
}
