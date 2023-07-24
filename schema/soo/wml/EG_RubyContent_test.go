package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestEG_RubyContentConstructor(t *testing.T) {
	v := wml.NewEG_RubyContent()
	if v == nil {
		t.Errorf("wml.NewEG_RubyContent must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.EG_RubyContent should validate: %s", err)
	}
}

func TestEG_RubyContentMarshalUnmarshal(t *testing.T) {
	v := wml.NewEG_RubyContent()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewEG_RubyContent()
	xml.Unmarshal(buf, v2)
}
