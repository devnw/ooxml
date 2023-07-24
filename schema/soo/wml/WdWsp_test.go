package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestWdWspConstructor(t *testing.T) {
	v := wml.NewWdWsp()
	if v == nil {
		t.Errorf("wml.NewWdWsp must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.WdWsp should validate: %s", err)
	}
}

func TestWdWspMarshalUnmarshal(t *testing.T) {
	v := wml.NewWdWsp()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewWdWsp()
	xml.Unmarshal(buf, v2)
}
