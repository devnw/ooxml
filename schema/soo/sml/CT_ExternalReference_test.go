package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ExternalReferenceConstructor(t *testing.T) {
	v := sml.NewCT_ExternalReference()
	if v == nil {
		t.Errorf("sml.NewCT_ExternalReference must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_ExternalReference should validate: %s", err)
	}
}

func TestCT_ExternalReferenceMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_ExternalReference()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_ExternalReference()
	xml.Unmarshal(buf, v2)
}
