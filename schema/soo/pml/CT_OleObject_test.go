package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_OleObjectConstructor(t *testing.T) {
	v := pml.NewCT_OleObject()
	if v == nil {
		t.Errorf("pml.NewCT_OleObject must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_OleObject should validate: %s", err)
	}
}

func TestCT_OleObjectMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_OleObject()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_OleObject()
	xml.Unmarshal(buf, v2)
}
