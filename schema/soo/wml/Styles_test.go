package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestStylesConstructor(t *testing.T) {
	v := wml.NewStyles()
	if v == nil {
		t.Errorf("wml.NewStyles must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.Styles should validate: %s", err)
	}
}

func TestStylesMarshalUnmarshal(t *testing.T) {
	v := wml.NewStyles()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewStyles()
	xml.Unmarshal(buf, v2)
}
