package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_XmlCellPrConstructor(t *testing.T) {
	v := sml.NewCT_XmlCellPr()
	if v == nil {
		t.Errorf("sml.NewCT_XmlCellPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_XmlCellPr should validate: %s", err)
	}
}

func TestCT_XmlCellPrMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_XmlCellPr()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_XmlCellPr()
	xml.Unmarshal(buf, v2)
}
