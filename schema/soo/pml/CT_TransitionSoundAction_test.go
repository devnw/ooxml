package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TransitionSoundActionConstructor(t *testing.T) {
	v := pml.NewCT_TransitionSoundAction()
	if v == nil {
		t.Errorf("pml.NewCT_TransitionSoundAction must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TransitionSoundAction should validate: %s", err)
	}
}

func TestCT_TransitionSoundActionMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TransitionSoundAction()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TransitionSoundAction()
	xml.Unmarshal(buf, v2)
}
