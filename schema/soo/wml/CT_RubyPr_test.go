package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_RubyPrConstructor(t *testing.T) {
	v := wml.NewCT_RubyPr()
	if v == nil {
		t.Errorf("wml.NewCT_RubyPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_RubyPr should validate: %s", err)
	}
}

func TestCT_RubyPrMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_RubyPr()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_RubyPr()
	xml.Unmarshal(buf, v2)
}
