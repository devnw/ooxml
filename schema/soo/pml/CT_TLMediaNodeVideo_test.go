package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TLMediaNodeVideoConstructor(t *testing.T) {
	v := pml.NewCT_TLMediaNodeVideo()
	if v == nil {
		t.Errorf("pml.NewCT_TLMediaNodeVideo must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLMediaNodeVideo should validate: %s", err)
	}
}

func TestCT_TLMediaNodeVideoMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLMediaNodeVideo()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLMediaNodeVideo()
	xml.Unmarshal(buf, v2)
}
