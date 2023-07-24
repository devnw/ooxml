package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_AlphaReplaceEffectConstructor(t *testing.T) {
	v := dml.NewCT_AlphaReplaceEffect()
	if v == nil {
		t.Errorf("dml.NewCT_AlphaReplaceEffect must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_AlphaReplaceEffect should validate: %s", err)
	}
}

func TestCT_AlphaReplaceEffectMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_AlphaReplaceEffect()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_AlphaReplaceEffect()
	xml.Unmarshal(buf, v2)
}
