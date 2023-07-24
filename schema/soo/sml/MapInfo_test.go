package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestMapInfoConstructor(t *testing.T) {
	v := sml.NewMapInfo()
	if v == nil {
		t.Errorf("sml.NewMapInfo must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.MapInfo should validate: %s", err)
	}
}

func TestMapInfoMarshalUnmarshal(t *testing.T) {
	v := sml.NewMapInfo()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewMapInfo()
	xml.Unmarshal(buf, v2)
}
