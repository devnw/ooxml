package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_SplitTransitionConstructor(t *testing.T) {
	v := pml.NewCT_SplitTransition()
	if v == nil {
		t.Errorf("pml.NewCT_SplitTransition must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_SplitTransition should validate: %s", err)
	}
}

func TestCT_SplitTransitionMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_SplitTransition()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_SplitTransition()
	xml.Unmarshal(buf, v2)
}
