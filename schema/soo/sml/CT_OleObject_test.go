package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_OleObjectConstructor(t *testing.T) {
	v := sml.NewCT_OleObject()
	if v == nil {
		t.Errorf("sml.NewCT_OleObject must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_OleObject should validate: %s", err)
	}
}

func TestCT_OleObjectMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_OleObject()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_OleObject()
	xml.Unmarshal(buf, v2)
}
