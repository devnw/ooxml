package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_ObjectEmbedConstructor(t *testing.T) {
	v := wml.NewCT_ObjectEmbed()
	if v == nil {
		t.Errorf("wml.NewCT_ObjectEmbed must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_ObjectEmbed should validate: %s", err)
	}
}

func TestCT_ObjectEmbedMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_ObjectEmbed()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_ObjectEmbed()
	xml.Unmarshal(buf, v2)
}
