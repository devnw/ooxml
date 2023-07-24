package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ExtensionConstructor(t *testing.T) {
	v := sml.NewCT_Extension()
	if v == nil {
		t.Errorf("sml.NewCT_Extension must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Extension should validate: %s", err)
	}
}

func TestCT_ExtensionMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Extension()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Extension()
	xml.Unmarshal(buf, v2)
}
