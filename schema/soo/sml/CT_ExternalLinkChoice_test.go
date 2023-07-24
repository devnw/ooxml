package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ExternalLinkChoiceConstructor(t *testing.T) {
	v := sml.NewCT_ExternalLinkChoice()
	if v == nil {
		t.Errorf("sml.NewCT_ExternalLinkChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_ExternalLinkChoice should validate: %s", err)
	}
}

func TestCT_ExternalLinkChoiceMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_ExternalLinkChoice()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_ExternalLinkChoice()
	xml.Unmarshal(buf, v2)
}
