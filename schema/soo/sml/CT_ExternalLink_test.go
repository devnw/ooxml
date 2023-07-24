package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ExternalLinkConstructor(t *testing.T) {
	v := sml.NewCT_ExternalLink()
	if v == nil {
		t.Errorf("sml.NewCT_ExternalLink must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_ExternalLink should validate: %s", err)
	}
}

func TestCT_ExternalLinkMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_ExternalLink()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_ExternalLink()
	xml.Unmarshal(buf, v2)
}
