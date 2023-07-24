package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_CustomXmlPrConstructor(t *testing.T) {
	v := wml.NewCT_CustomXmlPr()
	if v == nil {
		t.Errorf("wml.NewCT_CustomXmlPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_CustomXmlPr should validate: %s", err)
	}
}

func TestCT_CustomXmlPrMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_CustomXmlPr()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_CustomXmlPr()
	xml.Unmarshal(buf, v2)
}
