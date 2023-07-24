package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_EffectStyleListConstructor(t *testing.T) {
	v := dml.NewCT_EffectStyleList()
	if v == nil {
		t.Errorf("dml.NewCT_EffectStyleList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_EffectStyleList should validate: %s", err)
	}
}

func TestCT_EffectStyleListMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_EffectStyleList()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_EffectStyleList()
	xml.Unmarshal(buf, v2)
}
