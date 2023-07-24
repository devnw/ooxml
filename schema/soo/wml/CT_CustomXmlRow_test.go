package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_CustomXmlRowConstructor(t *testing.T) {
	v := wml.NewCT_CustomXmlRow()
	if v == nil {
		t.Errorf("wml.NewCT_CustomXmlRow must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_CustomXmlRow should validate: %s", err)
	}
}

func TestCT_CustomXmlRowMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_CustomXmlRow()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_CustomXmlRow()
	xml.Unmarshal(buf, v2)
}
