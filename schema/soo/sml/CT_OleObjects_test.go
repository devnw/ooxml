package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_OleObjectsConstructor(t *testing.T) {
	v := sml.NewCT_OleObjects()
	if v == nil {
		t.Errorf("sml.NewCT_OleObjects must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_OleObjects should validate: %s", err)
	}
}

func TestCT_OleObjectsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_OleObjects()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_OleObjects()
	xml.Unmarshal(buf, v2)
}
