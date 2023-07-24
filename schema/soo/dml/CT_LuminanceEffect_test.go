package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_LuminanceEffectConstructor(t *testing.T) {
	v := dml.NewCT_LuminanceEffect()
	if v == nil {
		t.Errorf("dml.NewCT_LuminanceEffect must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_LuminanceEffect should validate: %s", err)
	}
}

func TestCT_LuminanceEffectMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_LuminanceEffect()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_LuminanceEffect()
	xml.Unmarshal(buf, v2)
}
