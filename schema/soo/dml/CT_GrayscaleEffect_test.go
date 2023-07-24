package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_GrayscaleEffectConstructor(t *testing.T) {
	v := dml.NewCT_GrayscaleEffect()
	if v == nil {
		t.Errorf("dml.NewCT_GrayscaleEffect must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_GrayscaleEffect should validate: %s", err)
	}
}

func TestCT_GrayscaleEffectMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_GrayscaleEffect()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_GrayscaleEffect()
	xml.Unmarshal(buf, v2)
}
