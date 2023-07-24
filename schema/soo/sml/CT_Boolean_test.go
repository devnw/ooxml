package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_BooleanConstructor(t *testing.T) {
	v := sml.NewCT_Boolean()
	if v == nil {
		t.Errorf("sml.NewCT_Boolean must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Boolean should validate: %s", err)
	}
}

func TestCT_BooleanMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Boolean()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Boolean()
	xml.Unmarshal(buf, v2)
}
