package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_WebPublishObjectsConstructor(t *testing.T) {
	v := sml.NewCT_WebPublishObjects()
	if v == nil {
		t.Errorf("sml.NewCT_WebPublishObjects must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_WebPublishObjects should validate: %s", err)
	}
}

func TestCT_WebPublishObjectsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_WebPublishObjects()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_WebPublishObjects()
	xml.Unmarshal(buf, v2)
}
