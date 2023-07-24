package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_OrientationTransitionConstructor(t *testing.T) {
	v := pml.NewCT_OrientationTransition()
	if v == nil {
		t.Errorf("pml.NewCT_OrientationTransition must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_OrientationTransition should validate: %s", err)
	}
}

func TestCT_OrientationTransitionMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_OrientationTransition()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_OrientationTransition()
	xml.Unmarshal(buf, v2)
}
