package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_SingleXmlCellConstructor(t *testing.T) {
	v := sml.NewCT_SingleXmlCell()
	if v == nil {
		t.Errorf("sml.NewCT_SingleXmlCell must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_SingleXmlCell should validate: %s", err)
	}
}

func TestCT_SingleXmlCellMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_SingleXmlCell()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_SingleXmlCell()
	xml.Unmarshal(buf, v2)
}
