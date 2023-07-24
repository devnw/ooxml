package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_TransformEffectConstructor(t *testing.T) {
	v := dml.NewCT_TransformEffect()
	if v == nil {
		t.Errorf("dml.NewCT_TransformEffect must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TransformEffect should validate: %s", err)
	}
}

func TestCT_TransformEffectMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TransformEffect()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TransformEffect()
	xml.Unmarshal(buf, v2)
}
