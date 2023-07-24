package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_MapConstructor(t *testing.T) {
	v := sml.NewCT_Map()
	if v == nil {
		t.Errorf("sml.NewCT_Map must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Map should validate: %s", err)
	}
}

func TestCT_MapMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Map()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Map()
	xml.Unmarshal(buf, v2)
}
