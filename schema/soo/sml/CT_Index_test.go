package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_IndexConstructor(t *testing.T) {
	v := sml.NewCT_Index()
	if v == nil {
		t.Errorf("sml.NewCT_Index must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Index should validate: %s", err)
	}
}

func TestCT_IndexMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Index()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Index()
	xml.Unmarshal(buf, v2)
}
