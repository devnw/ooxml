package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ExternalReferencesConstructor(t *testing.T) {
	v := sml.NewCT_ExternalReferences()
	if v == nil {
		t.Errorf("sml.NewCT_ExternalReferences must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_ExternalReferences should validate: %s", err)
	}
}

func TestCT_ExternalReferencesMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_ExternalReferences()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_ExternalReferences()
	xml.Unmarshal(buf, v2)
}
