package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_StringTagConstructor(t *testing.T) {
	v := pml.NewCT_StringTag()
	if v == nil {
		t.Errorf("pml.NewCT_StringTag must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_StringTag should validate: %s", err)
	}
}

func TestCT_StringTagMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_StringTag()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_StringTag()
	xml.Unmarshal(buf, v2)
}
