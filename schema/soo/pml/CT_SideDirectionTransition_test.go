package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_SideDirectionTransitionConstructor(t *testing.T) {
	v := pml.NewCT_SideDirectionTransition()
	if v == nil {
		t.Errorf("pml.NewCT_SideDirectionTransition must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_SideDirectionTransition should validate: %s", err)
	}
}

func TestCT_SideDirectionTransitionMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_SideDirectionTransition()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_SideDirectionTransition()
	xml.Unmarshal(buf, v2)
}
