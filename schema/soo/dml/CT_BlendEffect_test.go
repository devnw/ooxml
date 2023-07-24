package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_BlendEffectConstructor(t *testing.T) {
	v := dml.NewCT_BlendEffect()
	if v == nil {
		t.Errorf("dml.NewCT_BlendEffect must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_BlendEffect should validate: %s", err)
	}
}

func TestCT_BlendEffectMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_BlendEffect()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_BlendEffect()
	xml.Unmarshal(buf, v2)
}
