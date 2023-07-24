package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_TextPrConstructor(t *testing.T) {
	v := sml.NewCT_TextPr()
	if v == nil {
		t.Errorf("sml.NewCT_TextPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_TextPr should validate: %s", err)
	}
}

func TestCT_TextPrMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_TextPr()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_TextPr()
	xml.Unmarshal(buf, v2)
}
