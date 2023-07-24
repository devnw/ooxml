package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_FFHelpTextConstructor(t *testing.T) {
	v := wml.NewCT_FFHelpText()
	if v == nil {
		t.Errorf("wml.NewCT_FFHelpText must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_FFHelpText should validate: %s", err)
	}
}

func TestCT_FFHelpTextMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_FFHelpText()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_FFHelpText()
	xml.Unmarshal(buf, v2)
}
