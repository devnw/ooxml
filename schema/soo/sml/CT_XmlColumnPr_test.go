package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_XmlColumnPrConstructor(t *testing.T) {
	v := sml.NewCT_XmlColumnPr()
	if v == nil {
		t.Errorf("sml.NewCT_XmlColumnPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_XmlColumnPr should validate: %s", err)
	}
}

func TestCT_XmlColumnPrMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_XmlColumnPr()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_XmlColumnPr()
	xml.Unmarshal(buf, v2)
}
