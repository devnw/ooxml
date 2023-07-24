package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_HyperlinksConstructor(t *testing.T) {
	v := sml.NewCT_Hyperlinks()
	if v == nil {
		t.Errorf("sml.NewCT_Hyperlinks must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Hyperlinks should validate: %s", err)
	}
}

func TestCT_HyperlinksMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Hyperlinks()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Hyperlinks()
	xml.Unmarshal(buf, v2)
}
