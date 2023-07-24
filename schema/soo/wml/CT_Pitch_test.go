package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_PitchConstructor(t *testing.T) {
	v := wml.NewCT_Pitch()
	if v == nil {
		t.Errorf("wml.NewCT_Pitch must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Pitch should validate: %s", err)
	}
}

func TestCT_PitchMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Pitch()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Pitch()
	xml.Unmarshal(buf, v2)
}
