package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_GuideConstructor(t *testing.T) {
	v := pml.NewCT_Guide()
	if v == nil {
		t.Errorf("pml.NewCT_Guide must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_Guide should validate: %s", err)
	}
}

func TestCT_GuideMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_Guide()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_Guide()
	xml.Unmarshal(buf, v2)
}
