package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestWdWgpConstructor(t *testing.T) {
	v := wml.NewWdWgp()
	if v == nil {
		t.Errorf("wml.NewWdWgp must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.WdWgp should validate: %s", err)
	}
}

func TestWdWgpMarshalUnmarshal(t *testing.T) {
	v := wml.NewWdWgp()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewWdWgp()
	xml.Unmarshal(buf, v2)
}
