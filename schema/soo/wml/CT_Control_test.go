package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_ControlConstructor(t *testing.T) {
	v := wml.NewCT_Control()
	if v == nil {
		t.Errorf("wml.NewCT_Control must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Control should validate: %s", err)
	}
}

func TestCT_ControlMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Control()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Control()
	xml.Unmarshal(buf, v2)
}
