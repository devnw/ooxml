package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_XmlPrConstructor(t *testing.T) {
	v := sml.NewCT_XmlPr()
	if v == nil {
		t.Errorf("sml.NewCT_XmlPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_XmlPr should validate: %s", err)
	}
}

func TestCT_XmlPrMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_XmlPr()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_XmlPr()
	xml.Unmarshal(buf, v2)
}
