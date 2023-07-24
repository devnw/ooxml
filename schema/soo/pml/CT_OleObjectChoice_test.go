package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_OleObjectChoiceConstructor(t *testing.T) {
	v := pml.NewCT_OleObjectChoice()
	if v == nil {
		t.Errorf("pml.NewCT_OleObjectChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_OleObjectChoice should validate: %s", err)
	}
}

func TestCT_OleObjectChoiceMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_OleObjectChoice()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_OleObjectChoice()
	xml.Unmarshal(buf, v2)
}
