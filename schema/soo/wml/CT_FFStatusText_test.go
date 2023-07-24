package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_FFStatusTextConstructor(t *testing.T) {
	v := wml.NewCT_FFStatusText()
	if v == nil {
		t.Errorf("wml.NewCT_FFStatusText must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_FFStatusText should validate: %s", err)
	}
}

func TestCT_FFStatusTextMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_FFStatusText()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_FFStatusText()
	xml.Unmarshal(buf, v2)
}
