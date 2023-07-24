package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_SelectionConstructor(t *testing.T) {
	v := sml.NewCT_Selection()
	if v == nil {
		t.Errorf("sml.NewCT_Selection must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Selection should validate: %s", err)
	}
}

func TestCT_SelectionMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Selection()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Selection()
	xml.Unmarshal(buf, v2)
}
