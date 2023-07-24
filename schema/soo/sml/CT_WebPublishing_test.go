package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_WebPublishingConstructor(t *testing.T) {
	v := sml.NewCT_WebPublishing()
	if v == nil {
		t.Errorf("sml.NewCT_WebPublishing must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_WebPublishing should validate: %s", err)
	}
}

func TestCT_WebPublishingMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_WebPublishing()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_WebPublishing()
	xml.Unmarshal(buf, v2)
}
