package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TLByAnimateColorTransformConstructor(t *testing.T) {
	v := pml.NewCT_TLByAnimateColorTransform()
	if v == nil {
		t.Errorf("pml.NewCT_TLByAnimateColorTransform must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLByAnimateColorTransform should validate: %s", err)
	}
}

func TestCT_TLByAnimateColorTransformMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLByAnimateColorTransform()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLByAnimateColorTransform()
	xml.Unmarshal(buf, v2)
}
