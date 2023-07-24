package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_AlphaFloorEffectConstructor(t *testing.T) {
	v := dml.NewCT_AlphaFloorEffect()
	if v == nil {
		t.Errorf("dml.NewCT_AlphaFloorEffect must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_AlphaFloorEffect should validate: %s", err)
	}
}

func TestCT_AlphaFloorEffectMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_AlphaFloorEffect()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_AlphaFloorEffect()
	xml.Unmarshal(buf, v2)
}
