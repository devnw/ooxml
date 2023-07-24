package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ControlsConstructor(t *testing.T) {
	v := sml.NewCT_Controls()
	if v == nil {
		t.Errorf("sml.NewCT_Controls must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Controls should validate: %s", err)
	}
}

func TestCT_ControlsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Controls()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Controls()
	xml.Unmarshal(buf, v2)
}
