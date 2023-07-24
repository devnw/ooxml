package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TLMediaNodeAudioConstructor(t *testing.T) {
	v := pml.NewCT_TLMediaNodeAudio()
	if v == nil {
		t.Errorf("pml.NewCT_TLMediaNodeAudio must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLMediaNodeAudio should validate: %s", err)
	}
}

func TestCT_TLMediaNodeAudioMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLMediaNodeAudio()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLMediaNodeAudio()
	xml.Unmarshal(buf, v2)
}
