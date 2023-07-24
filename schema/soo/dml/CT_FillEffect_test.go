package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_FillEffectConstructor(t *testing.T) {
	v := dml.NewCT_FillEffect()
	if v == nil {
		t.Errorf("dml.NewCT_FillEffect must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_FillEffect should validate: %s", err)
	}
}

func TestCT_FillEffectMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_FillEffect()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_FillEffect()
	xml.Unmarshal(buf, v2)
}
