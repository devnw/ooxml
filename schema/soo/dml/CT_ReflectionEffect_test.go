package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_ReflectionEffectConstructor(t *testing.T) {
	v := dml.NewCT_ReflectionEffect()
	if v == nil {
		t.Errorf("dml.NewCT_ReflectionEffect must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_ReflectionEffect should validate: %s", err)
	}
}

func TestCT_ReflectionEffectMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_ReflectionEffect()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_ReflectionEffect()
	xml.Unmarshal(buf, v2)
}
