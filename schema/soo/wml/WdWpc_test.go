package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestWdWpcConstructor(t *testing.T) {
	v := wml.NewWdWpc()
	if v == nil {
		t.Errorf("wml.NewWdWpc must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.WdWpc should validate: %s", err)
	}
}

func TestWdWpcMarshalUnmarshal(t *testing.T) {
	v := wml.NewWdWpc()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewWdWpc()
	xml.Unmarshal(buf, v2)
}
