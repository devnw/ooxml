package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_CustomXmlCellConstructor(t *testing.T) {
	v := wml.NewCT_CustomXmlCell()
	if v == nil {
		t.Errorf("wml.NewCT_CustomXmlCell must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_CustomXmlCell should validate: %s", err)
	}
}

func TestCT_CustomXmlCellMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_CustomXmlCell()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_CustomXmlCell()
	xml.Unmarshal(buf, v2)
}
