package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_BorderPrConstructor(t *testing.T) {
	v := sml.NewCT_BorderPr()
	if v == nil {
		t.Errorf("sml.NewCT_BorderPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_BorderPr should validate: %s", err)
	}
}

func TestCT_BorderPrMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_BorderPr()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_BorderPr()
	xml.Unmarshal(buf, v2)
}
