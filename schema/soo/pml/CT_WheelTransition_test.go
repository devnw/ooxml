package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_WheelTransitionConstructor(t *testing.T) {
	v := pml.NewCT_WheelTransition()
	if v == nil {
		t.Errorf("pml.NewCT_WheelTransition must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_WheelTransition should validate: %s", err)
	}
}

func TestCT_WheelTransitionMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_WheelTransition()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_WheelTransition()
	xml.Unmarshal(buf, v2)
}
