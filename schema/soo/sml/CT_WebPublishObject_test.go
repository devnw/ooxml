package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_WebPublishObjectConstructor(t *testing.T) {
	v := sml.NewCT_WebPublishObject()
	if v == nil {
		t.Errorf("sml.NewCT_WebPublishObject must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_WebPublishObject should validate: %s", err)
	}
}

func TestCT_WebPublishObjectMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_WebPublishObject()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_WebPublishObject()
	xml.Unmarshal(buf, v2)
}
