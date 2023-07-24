package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_AlphaModulateFixedEffectConstructor(t *testing.T) {
	v := dml.NewCT_AlphaModulateFixedEffect()
	if v == nil {
		t.Errorf("dml.NewCT_AlphaModulateFixedEffect must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_AlphaModulateFixedEffect should validate: %s", err)
	}
}

func TestCT_AlphaModulateFixedEffectMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_AlphaModulateFixedEffect()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_AlphaModulateFixedEffect()
	xml.Unmarshal(buf, v2)
}
