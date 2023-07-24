package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_RubyAlignConstructor(t *testing.T) {
	v := wml.NewCT_RubyAlign()
	if v == nil {
		t.Errorf("wml.NewCT_RubyAlign must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_RubyAlign should validate: %s", err)
	}
}

func TestCT_RubyAlignMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_RubyAlign()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_RubyAlign()
	xml.Unmarshal(buf, v2)
}
