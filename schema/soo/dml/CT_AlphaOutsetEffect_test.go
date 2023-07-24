package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_AlphaOutsetEffectConstructor(t *testing.T) {
	v := dml.NewCT_AlphaOutsetEffect()
	if v == nil {
		t.Errorf("dml.NewCT_AlphaOutsetEffect must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_AlphaOutsetEffect should validate: %s", err)
	}
}

func TestCT_AlphaOutsetEffectMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_AlphaOutsetEffect()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_AlphaOutsetEffect()
	xml.Unmarshal(buf, v2)
}
