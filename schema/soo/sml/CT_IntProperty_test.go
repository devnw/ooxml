package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_IntPropertyConstructor(t *testing.T) {
	v := sml.NewCT_IntProperty()
	if v == nil {
		t.Errorf("sml.NewCT_IntProperty must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_IntProperty should validate: %s", err)
	}
}

func TestCT_IntPropertyMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_IntProperty()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_IntProperty()
	xml.Unmarshal(buf, v2)
}
