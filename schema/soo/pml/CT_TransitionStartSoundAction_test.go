package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TransitionStartSoundActionConstructor(t *testing.T) {
	v := pml.NewCT_TransitionStartSoundAction()
	if v == nil {
		t.Errorf("pml.NewCT_TransitionStartSoundAction must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TransitionStartSoundAction should validate: %s", err)
	}
}

func TestCT_TransitionStartSoundActionMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TransitionStartSoundAction()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TransitionStartSoundAction()
	xml.Unmarshal(buf, v2)
}
