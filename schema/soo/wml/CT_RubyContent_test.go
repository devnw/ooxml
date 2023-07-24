package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_RubyContentConstructor(t *testing.T) {
	v := wml.NewCT_RubyContent()
	if v == nil {
		t.Errorf("wml.NewCT_RubyContent must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_RubyContent should validate: %s", err)
	}
}

func TestCT_RubyContentMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_RubyContent()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_RubyContent()
	xml.Unmarshal(buf, v2)
}
