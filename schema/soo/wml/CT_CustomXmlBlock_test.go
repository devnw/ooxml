package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_CustomXmlBlockConstructor(t *testing.T) {
	v := wml.NewCT_CustomXmlBlock()
	if v == nil {
		t.Errorf("wml.NewCT_CustomXmlBlock must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_CustomXmlBlock should validate: %s", err)
	}
}

func TestCT_CustomXmlBlockMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_CustomXmlBlock()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_CustomXmlBlock()
	xml.Unmarshal(buf, v2)
}
