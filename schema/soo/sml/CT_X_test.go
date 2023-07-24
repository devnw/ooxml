package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_XConstructor(t *testing.T) {
	v := sml.NewCT_X()
	if v == nil {
		t.Errorf("sml.NewCT_X must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_X should validate: %s", err)
	}
}

func TestCT_XMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_X()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_X()
	xml.Unmarshal(buf, v2)
}
