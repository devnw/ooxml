package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_AlphaModulateEffectConstructor(t *testing.T) {
	v := dml.NewCT_AlphaModulateEffect()
	if v == nil {
		t.Errorf("dml.NewCT_AlphaModulateEffect must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_AlphaModulateEffect should validate: %s", err)
	}
}

func TestCT_AlphaModulateEffectMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_AlphaModulateEffect()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_AlphaModulateEffect()
	xml.Unmarshal(buf, v2)
}
