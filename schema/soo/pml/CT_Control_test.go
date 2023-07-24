package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_ControlConstructor(t *testing.T) {
	v := pml.NewCT_Control()
	if v == nil {
		t.Errorf("pml.NewCT_Control must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_Control should validate: %s", err)
	}
}

func TestCT_ControlMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_Control()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_Control()
	xml.Unmarshal(buf, v2)
}
