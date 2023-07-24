package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_CustomXmlRunConstructor(t *testing.T) {
	v := wml.NewCT_CustomXmlRun()
	if v == nil {
		t.Errorf("wml.NewCT_CustomXmlRun must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_CustomXmlRun should validate: %s", err)
	}
}

func TestCT_CustomXmlRunMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_CustomXmlRun()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_CustomXmlRun()
	xml.Unmarshal(buf, v2)
}
