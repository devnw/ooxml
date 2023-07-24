package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestWdAnchorConstructor(t *testing.T) {
	v := wml.NewWdAnchor()
	if v == nil {
		t.Errorf("wml.NewWdAnchor must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.WdAnchor should validate: %s", err)
	}
}

func TestWdAnchorMarshalUnmarshal(t *testing.T) {
	v := wml.NewWdAnchor()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewWdAnchor()
	xml.Unmarshal(buf, v2)
}
