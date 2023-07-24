package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcClippathConstructor(t *testing.T) {
	v := vml.NewOfcClippath()
	if v == nil {
		t.Errorf("vml.NewOfcClippath must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcClippath should validate: %s", err)
	}
}

func TestOfcClippathMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcClippath()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcClippath()
	xml.Unmarshal(buf, v2)
}
