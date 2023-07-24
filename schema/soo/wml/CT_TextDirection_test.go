package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_TextDirectionConstructor(t *testing.T) {
	v := wml.NewCT_TextDirection()
	if v == nil {
		t.Errorf("wml.NewCT_TextDirection must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_TextDirection should validate: %s", err)
	}
}

func TestCT_TextDirectionMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_TextDirection()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_TextDirection()
	xml.Unmarshal(buf, v2)
}
