package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ObjectPrConstructor(t *testing.T) {
	v := sml.NewCT_ObjectPr()
	if v == nil {
		t.Errorf("sml.NewCT_ObjectPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_ObjectPr should validate: %s", err)
	}
}

func TestCT_ObjectPrMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_ObjectPr()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_ObjectPr()
	xml.Unmarshal(buf, v2)
}
