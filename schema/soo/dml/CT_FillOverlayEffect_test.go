package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_FillOverlayEffectConstructor(t *testing.T) {
	v := dml.NewCT_FillOverlayEffect()
	if v == nil {
		t.Errorf("dml.NewCT_FillOverlayEffect must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_FillOverlayEffect should validate: %s", err)
	}
}

func TestCT_FillOverlayEffectMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_FillOverlayEffect()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_FillOverlayEffect()
	xml.Unmarshal(buf, v2)
}
