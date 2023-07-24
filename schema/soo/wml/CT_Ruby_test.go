package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_RubyConstructor(t *testing.T) {
	v := wml.NewCT_Ruby()
	if v == nil {
		t.Errorf("wml.NewCT_Ruby must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Ruby should validate: %s", err)
	}
}

func TestCT_RubyMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Ruby()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Ruby()
	xml.Unmarshal(buf, v2)
}
