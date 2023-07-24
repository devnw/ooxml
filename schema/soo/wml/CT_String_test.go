package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_StringConstructor(t *testing.T) {
	v := wml.NewCT_String()
	if v == nil {
		t.Errorf("wml.NewCT_String must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_String should validate: %s", err)
	}
}

func TestCT_StringMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_String()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_String()
	xml.Unmarshal(buf, v2)
}
