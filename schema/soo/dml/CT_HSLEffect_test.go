package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_HSLEffectConstructor(t *testing.T) {
	v := dml.NewCT_HSLEffect()
	if v == nil {
		t.Errorf("dml.NewCT_HSLEffect must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_HSLEffect should validate: %s", err)
	}
}

func TestCT_HSLEffectMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_HSLEffect()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_HSLEffect()
	xml.Unmarshal(buf, v2)
}
