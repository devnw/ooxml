package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_TintEffectConstructor(t *testing.T) {
	v := dml.NewCT_TintEffect()
	if v == nil {
		t.Errorf("dml.NewCT_TintEffect must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TintEffect should validate: %s", err)
	}
}

func TestCT_TintEffectMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TintEffect()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TintEffect()
	xml.Unmarshal(buf, v2)
}
