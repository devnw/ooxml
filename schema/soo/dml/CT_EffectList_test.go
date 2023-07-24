package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_EffectListConstructor(t *testing.T) {
	v := dml.NewCT_EffectList()
	if v == nil {
		t.Errorf("dml.NewCT_EffectList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_EffectList should validate: %s", err)
	}
}

func TestCT_EffectListMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_EffectList()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_EffectList()
	xml.Unmarshal(buf, v2)
}
